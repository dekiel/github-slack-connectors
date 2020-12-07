import os
from slack_sdk import WebClient
from slack_sdk.errors import SlackApiError


def main(event, context):
	client = WebClient(base_url=os.environ['KYMA_SLACK_GATEWAY_URL'])
	label = event["data"]["label"]["name"]
	title = event["data"]["issue"]["title"]
	try:
		assignee = "Issue is assigned to `{}`.".format(event["data"]["issue"]["assignee"]["login"])
	except TypeError:
		assignee = "Issue is not assigned."
	sender = event["data"]["sender"]["login"]
	issue_url = event["data"]["issue"]["html_url"]
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
												"type": "section",
												"text":
													{
														"type": "mrkdwn",
														"text": "`{}` labeled issue `{}` as `{}`.\n{} <{}|You can find it here>".format(sender, title, label, assignee, issue_url)
													}
											},
											{
												"type": "section",
												"text":
													{
														"type": "mrkdwn",
														"text": "Issue is assigned to *{}*. <{}|You can find it here>".format(
															assignee, issue_url)
													}
												},
											])
		assert response["ok"]
	except SlackApiError as e:
		# You will get a SlackApiError if "ok" is False
		assert e.response["ok"] is False
		assert e.response["error"]  # str like 'invalid_auth', 'channel_not_found'
		print(f"Got an error: {e.response['error']}")
