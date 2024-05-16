import argparse
import os
from github import Github, InputGitAuthor

github_token = os.getenv('TOKEN')
repository_source = os.getenv('REPOSITORY_SOURCE')

def create_new_tag(github_token, repo_name):
    g = Github(github_token)
    repo = g.get_repo(repo_name)

    tags = repo.get_tags()
    latest_tag = tags[0].name

    tag_without_prefix = latest_tag.lstrip('v')

    major, minor, patch = map(int, tag_without_prefix.split('.'))
    new_tag = f"v{major}.{minor}.{patch + 1}"

    repo.create_git_tag(new_tag, f"Version {new_tag}", repo.get_commits()[0].sha, "commit", tagger=InputGitAuthor("Automated Tagging", "noreply@github.com"))

    repo.create_git_ref(f"refs/tags/{new_tag}", repo.get_commits()[0].sha)

    print(f"Nova tag {new_tag} criada com sucesso.")

if __name__ == '__main__':
    print(repository_source)
    create_new_tag(github_token, repository_source)
