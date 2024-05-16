import argparse
from github import Github, InputGitAuthor

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
    parser = argparse.ArgumentParser(description='Cria e envia uma nova tag para um repositório no GitHub.')
    parser.add_argument('--github-token', required=True, help='GitHub token para autenticação')
    parser.add_argument('--repository', required=True, help='Nome do repositório (por exemplo, "nomeusuario/nomerepositorio")')
    
    args = parser.parse_args()

    create_new_tag(args.github_token, args.repository)
