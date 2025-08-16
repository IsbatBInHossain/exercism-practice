import 'dotenv/config';
import { Octokit } from 'octokit';

const {
  GH_TOKEN, // pass via workflow env
  GITHUB_REPOSITORY, // "owner/repo" (Actions provides)
  PR_NUMBER, // pull number if event-driven; else undefined
  AUTO_MERGE = 'false', // "true" to auto-merge clean PRs
  MERGE_METHOD = 'squash', // "merge" | "squash" | "rebase"
} = process.env;

if (!GH_TOKEN) {
  console.error('Missing GH_TOKEN env');
  process.exit(1);
}
const octokit = new Octokit({ auth: GH_TOKEN });

const sleep = ms => new Promise(r => setTimeout(r, ms));
const marker = '<!-- pr-mergeability-bot -->';

function parseRepo(repoStr) {
  const [owner, repo] = (repoStr ?? '').split('/');
  if (!owner || !repo) {
    console.error('GITHUB_REPOSITORY must be like owner/repo');
    process.exit(1);
  }
  return { owner, repo };
}

async function fetchPRWithMergeability(owner, repo, pull_number, retries = 8) {
  // mergeable can be null while GitHub computes it. Retry a few times.
  // See REST docs for PRs/mergeability. :contentReference[oaicite:1]{index=1}
  for (let i = 0; i < retries; i++) {
    const { data } = await octokit.rest.pulls.get({ owner, repo, pull_number });
    if (data.mergeable !== null) return data;
    await sleep(1500);
  }
  // last attempt
  const { data } = await octokit.rest.pulls.get({ owner, repo, pull_number });
  return data;
}

async function upsertComment(owner, repo, issue_number, body) {
  // Look for an existing bot comment and update it to avoid spam.
  const comments = await octokit.paginate(octokit.rest.issues.listComments, {
    owner,
    repo,
    issue_number,
    per_page: 100,
  });
  const existing = comments.find(c => (c.body || '').includes(marker));
  const finalBody = `${marker}\n${body}`;
  if (existing) {
    await octokit.rest.issues.updateComment({
      owner,
      repo,
      comment_id: existing.id,
      body: finalBody,
    });
  } else {
    await octokit.rest.issues.createComment({
      owner,
      repo,
      issue_number,
      body: finalBody,
    });
  }
}

async function maybeMerge(owner, repo, pull_number) {
  try {
    await octokit.rest.pulls.merge({
      owner,
      repo,
      pull_number,
      merge_method: MERGE_METHOD,
    });
    return true;
  } catch (e) {
    console.warn(`Auto-merge failed for #${pull_number}: ${e?.message}`);
    return false;
  }
}

function humanState(mergeable, mergeable_state) {
  if (mergeable === true) {
    return `✅ No conflicts (mergeable_state: \`${mergeable_state}\`)`;
  }
  if (mergeable === false) {
    return `❌ Has conflicts (mergeable_state: \`${mergeable_state}\`)`;
  }
  return `⏳ Mergeability still computing (mergeable=null)`;
}

async function processPR(owner, repo, pull_number) {
  const pr = await fetchPRWithMergeability(owner, repo, pull_number);
  const { mergeable, mergeable_state, html_url, title } = pr;

  const statusLine = humanState(mergeable, mergeable_state);
  let body = `**PR:** [#${pull_number}](${html_url}) — ${title}\n\n${statusLine}`;

  if (AUTO_MERGE === 'true' && mergeable === true) {
    const merged = await maybeMerge(owner, repo, pull_number);
    body += merged
      ? `\n\n🟢 Auto-merged with \`${MERGE_METHOD}\`.`
      : `\n\n🟡 Auto-merge attempted but failed.`;
  }

  await upsertComment(owner, repo, pull_number, body);
  console.log(`Processed PR #${pull_number}: ${statusLine}`);
}

async function main() {
  const { owner, repo } = parseRepo(GITHUB_REPOSITORY);
  if (PR_NUMBER) {
    await processPR(owner, repo, Number(PR_NUMBER));
    return;
  }
  // No PR_NUMBER: list all open PRs (for cron / manual runs)
  const prs = await octokit.paginate(octokit.rest.pulls.list, {
    owner,
    repo,
    state: 'open',
    per_page: 100,
  });
  if (prs.length === 0) {
    console.log('No open PRs');
    return;
  }
  for (const pr of prs) {
    await processPR(owner, repo, pr.number);
  }
}

main().catch(e => {
  console.error(e);
  process.exit(1);
});
