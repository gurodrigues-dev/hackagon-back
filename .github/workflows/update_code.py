import argparse
import re
import time
from github import Github

def update_docker_compose(content, old_version, new_version):
    pattern = re.compile(f'hackagon-back:{old_version}')
    updated_content = pattern.sub(f'hackagon-back:{new_version}', content)
    return updated_content

def main(github_token, repo_source, repo_receiver):
    g = Github(github_token)
    back_repo = g.get_repo(repo_source)
    infra_repo = g.get_repo(repo_receiver)

    tags = sorted(back_repo.get_tags(), key=lambda t: t.name, reverse=True)
    latest_tag = tags[0].name
    previous_tag = tags[1].name
    print(latest_tag, previous_tag)

    main_branch = infra_repo.get_branch('master')
    docker_compose_file = infra_repo.get_contents('docker-compose.yml', ref=main_branch.commit.sha)
    docker_compose_content = docker_compose_file.decoded_content.decode()

    new_docker_compose_content = update_docker_compose(docker_compose_content, previous_tag, latest_tag)

    new_branch_name = f'hackagon-deploy-{latest_tag}'
    main_sha = main_branch.commit.sha
    infra_repo.create_git_ref(ref=f'refs/heads/{new_branch_name}', sha=main_sha)

    infra_repo.update_file(docker_compose_file.path, f'Update hackagon-back to {latest_tag}', new_docker_compose_content, docker_compose_file.sha, branch=new_branch_name)

    pr = infra_repo.create_pull(
        title=f'Update hackagon-back to {latest_tag}',
        body=f'This PR updates the hackagon-back service to version {latest_tag}.',
        head=new_branch_name,
        base='master'
    )

    pr.merge()

    print(f'Successfully updated hackagon-back to {latest_tag} and merged the pull request.')

if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Update docker-compose with the latest hackagon-back version.')
    parser.add_argument('--github-token', required=True, help='GitHub token for authentication')
    parser.add_argument('--source', required=True, help='Source repository (e.g., gurodriguesdev/hackagon-back)')
    parser.add_argument('--receiver', required=True, help='Receiver repository (e.g., gurodriguesdev/hackagon-infra)')
    
    args = parser.parse_args()
    
    main(args.github_token, args.repository_source, args.repository_receiver)