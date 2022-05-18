import boto3
import subprocess

# TODO: User docker-py instead of subproc calls

# build image
# tag image
# get ECR credentials
# push image
# update service
# force new deployment

print("Starting deployment...")

#  aws ecr get-login-password --region ca-central-1 | docker login --username AWS --password-stdin 495679467660.dkr.ecr.ca-central-1.amazonaws.com
print("Getting ECR login password...")
get_password_proc = subprocess.Popen(['aws', 'ecr', 'get-login-password'],
                                     stdout=subprocess.PIPE,
                                     stderr=subprocess.PIPE,
                                     shell=True)
print("Done")

print("Logging into Docker...")
login_proc = subprocess.Popen(["docker", "login", "--username", "AWS", "--password-stdin", "495679467660.dkr.ecr.ca-central-1.amazonaws.com"],
                              stdin=get_password_proc.stdout,
                              stdout=subprocess.PIPE,
                              stderr=subprocess.PIPE,
                              shell=True)
login_stdout, stderr = login_proc.communicate()
if "Login Succeeded" not in str(login_stdout):
    print("Failed to login to docker")
    exit(1)
print("Done")

print("Building Docker image...")
docker_build_proc = subprocess.Popen(["docker", "build", "-t", "wonderfuldocker", "."],
                                     stdout=subprocess.PIPE,
                                     stderr=subprocess.PIPE,
                                     shell=True)
docker_build_stdout, stderr = docker_build_proc.communicate()
if docker_build_proc.returncode != 0:
    print("Failed to build docker image")
    exit(1)
print("Done")

print("Tagging Docker image...")
docker_tag_proc = subprocess.Popen(["docker", "tag", "wonderfuldocker:latest", "495679467660.dkr.ecr.ca-central-1.amazonaws.com/wonderfuldocker:latest"],
                                   stdout=subprocess.PIPE,
                                   stderr=subprocess.PIPE,
                                   shell=True)
docker_tag_stdout, stderr = docker_tag_proc.communicate()
if docker_tag_proc.returncode != 0:
    print("Failed to tag docker image")
    exit(1)
print("Done")

print("Pushing Docker image...")
docker_push_proc = subprocess.Popen(["docker", "push", "495679467660.dkr.ecr.ca-central-1.amazonaws.com/wonderfuldocker:latest"],
                                    stdout=subprocess.PIPE,
                                    stderr=subprocess.PIPE,
                                    shell=True)
docker_push_stdout, stderr = docker_push_proc.communicate()
if docker_push_proc.returncode != 0:
    print("Failed to push docker image")
    exit(1)
print("Done")

ecs_client = boto3.client('ecs')
cluster = "default"
service = "wonderfulestate"
print("Updating ECS service...")
ecs_client.update_service(forceNewDeployment=True, cluster=cluster,
                          service=service)
print("Done")