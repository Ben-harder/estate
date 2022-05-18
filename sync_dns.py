import boto3

print("Syncing Route53 entry to ECS task IP...")

ecs_client = boto3.client('ecs')
task = ecs_client.list_tasks(serviceName="wonderfulestate")
task_arn = task["taskArns"][0]
task_info = ecs_client.describe_tasks(tasks=[task_arn])
attachments = task_info["tasks"][0]["attachments"][0]["details"]
network_interface_id = ""
for attachment in attachments:
    if attachment["name"] == "networkInterfaceId":
        network_interface_id = attachment["value"]
        break

ec2_client = boto3.client('ec2')
network_interface_info = ec2_client.describe_network_interfaces(
    NetworkInterfaceIds=[network_interface_id])["NetworkInterfaces"][0]
public_ip = network_interface_info["Association"]["PublicIp"]
print("Found public IP ", public_ip)

route53_client = boto3.client('route53')
HOSTED_ZONE_ID = "Z074398527L3W0TLX5BTC"
route53_client.change_resource_record_sets(
    HostedZoneId=HOSTED_ZONE_ID,
    ChangeBatch={
        'Changes': [
            {
                'Action': 'UPSERT',
                'ResourceRecordSet': {
                    'Name': 'wonderfulestate.ca',
                    'Type': 'A',
                    'TTL': 300,
                    'ResourceRecords': [
                        {
                            'Value': public_ip
                        },
                    ],
                }
            }
        ]

    }
)

print("Done")