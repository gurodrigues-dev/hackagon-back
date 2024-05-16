import argparse
import os
from github import Github, InputGitAuthor

python_key = os.getenv('KEYPYTHON')
repository_source = os.getenv('REPOSITORY_SOURCE')

def decrypt(text):
    text_inverse = text[::-1]
    return text_inverse

def create_new_tag(python_key, repo_name):
    g = Github(python_key)
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
    token = decrypt(python_key)
    print(token)
    create_new_tag(token, repository_source)
