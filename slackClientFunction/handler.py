import os
from slack_sdk import WebClient
from slack_sdk.errors import SlackApiError


def main(event, context):
	client = WebClient(base_url=os.environ['KYMA_SLACK_CONNECTOR_E39A0FC6_C29C_4156_A205_DE7B24D4D480_GATEWAY_URL'])
	label = event["data"]["label"]["name"]
	title = event["data"]["issue"]["title"]
	number = event["data"]["issue"]["number"]
	repo = event["data"]["repository"]["name"]
	try:
		assignee = "Issue {} in repository {} is assigned to `{}`.".format(number, repo, event["data"]["issue"]["assignee"]["login"])
	except TypeError:
		assignee = "Issue {} in repository {} is not assigned.".format(number, repo)
	sender = event["data"]["sender"]["login"]
	issue_url = event["data"]["issue"]["html_url"]
	if (label == "internal-incident") or (label == "customer-incident"):
		try:
			response = client.chat_postMessage(channel='kyma-prow-dev-null',
											   blocks=[
												{
													"type": "context",
													"elements":
														[
															{
																"type": "image",
																"image_url": "https://mpng.subpng.com/20180802/bfy/kisspng-portable-network-graphics-computer-icons-clip-art-caribbean-blue-tag-icon-free-caribbean-blue-pric-5b63afe8224040.3966331515332597521403.jpg",
																"alt_text": "label"
															},
															{
																"type": "mrkdwn",
																"text": "SAP Github issue labeled"
															}
														]
												},
												{
													"type": "header",
													"text": {
														"type": "plain_text",
														"text": "SAP Github {}".format(label)
														}
												},
												{
													"type": "section",
													"text":
														{
															"type": "mrkdwn",
															"text": "*{}* labeled issue `{}` as `{}`.\n{} <{}|Check issue here.>".format(sender, title, label, assignee, issue_url)
														}
												},
												])
			assert response["ok"]
		except SlackApiError as e:
			# You will get a SlackApiError if "ok" is False
			assert e.response["ok"] is False
			assert e.response["error"]  # str like 'invalid_auth', 'channel_not_found'
			print(f"Got an error: {e.response['error']}")
