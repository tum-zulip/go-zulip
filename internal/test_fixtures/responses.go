// Package testfixtures provides mock API responses for benchmarking and testing.
package testfixtures

// Event Queue Responses

// RegisterQueueResponse is a minimal mock response for POST /api/v1/register.
const RegisterQueueResponse = `{
  "result": "success",
  "msg": "",
  "queue_id": "test-queue-1234567890",
  "last_event_id": 0,
  "zulip_feature_level": 200,
  "zulip_version": "8.0-dev",
  "zulip_merge_base": "8.0-dev"
}`

// EventsResponseEmpty is a mock response for GET /api/v1/events with no events.
const EventsResponseEmpty = `{
  "result": "success",
  "msg": "",
  "events": []
}`

// EventsResponseHeartbeat is a mock response with a heartbeat event.
const EventsResponseHeartbeat = `{
  "result": "success",
  "msg": "",
  "events": [
    {
      "type": "heartbeat",
      "id": 1
    }
  ]
}`

// EventsResponseMessage is a mock response with a message event.
const EventsResponseMessage = `{
  "result": "success",
  "msg": "",
  "events": [
    {
      "type": "message",
      "id": 2,
      "message": {
        "id": 12345,
        "sender_id": 100,
        "content": "Hello, world!",
        "timestamp": 1234567890,
        "client": "test",
        "subject": "test topic",
        "topic_links": [],
        "is_me_message": false,
        "reactions": [],
        "submessages": [],
        "sender_full_name": "Test User",
        "sender_email": "test@example.com",
        "sender_realm_str": "test",
        "display_recipient": "test stream",
        "type": "stream",
        "stream_id": 1,
        "avatar_url": null,
        "content_type": "text/html"
      }
    }
  ]
}`

// EventsResponsePresence is a mock response with a presence event.
const EventsResponsePresence = `{
  "result": "success",
  "msg": "",
  "events": [
    {
      "type": "presence",
      "id": 3,
      "user_id": 100,
      "server_timestamp": 1234567890,
      "presence": {
        "website": {
          "status": "active",
          "timestamp": 1234567890
        }
      }
    }
  ]
}`

// EventsResponseMultiple is a mock response with multiple events.
const EventsResponseMultiple = `{
  "result": "success",
  "msg": "",
  "events": [
    {
      "type": "heartbeat",
      "id": 1
    },
    {
      "type": "message",
      "id": 2,
      "message": {
        "id": 12345,
        "sender_id": 100,
        "content": "Hello!",
        "timestamp": 1234567890,
        "client": "test",
        "subject": "test",
        "topic_links": [],
        "is_me_message": false,
        "reactions": [],
        "submessages": [],
        "sender_full_name": "Test User",
        "sender_email": "test@example.com",
        "sender_realm_str": "test",
        "display_recipient": "test",
        "type": "stream",
        "stream_id": 1,
        "avatar_url": null,
        "content_type": "text/html"
      }
    },
    {
      "type": "heartbeat",
      "id": 3
    }
  ]
}`

// Regular API Responses

// SendMessageResponse is a mock response for POST /api/v1/messages.
const SendMessageResponse = `{
  "result": "success",
  "msg": "",
  "id": 42,
  "automatic_new_visibility_policy": 2,
  "deliver_at": null
}`

// GetChannelsResponse is a mock response for GET /api/v1/streams.
const GetChannelsResponse = `{
  "result": "success",
  "msg": "",
  "streams": [
    {
      "stream_id": 1,
      "name": "general",
      "description": "General discussion",
      "date_created": 1234567890,
      "invite_only": false,
      "rendered_description": "<p>General discussion</p>",
      "is_web_public": false,
      "stream_post_policy": 1,
      "history_public_to_subscribers": true,
      "first_message_id": 1,
      "message_retention_days": null,
      "is_announcement_only": false,
      "can_remove_subscribers_group": 2
    },
    {
      "stream_id": 2,
      "name": "random",
      "description": "Random stuff",
      "date_created": 1234567891,
      "invite_only": false,
      "rendered_description": "<p>Random stuff</p>",
      "is_web_public": false,
      "stream_post_policy": 1,
      "history_public_to_subscribers": true,
      "first_message_id": 100,
      "message_retention_days": null,
      "is_announcement_only": false,
      "can_remove_subscribers_group": 2
    },
    {
      "stream_id": 3,
      "name": "development",
      "description": "Development discussion",
      "date_created": 1234567892,
      "invite_only": true,
      "rendered_description": "<p>Development discussion</p>",
      "is_web_public": false,
      "stream_post_policy": 1,
      "history_public_to_subscribers": false,
      "first_message_id": 200,
      "message_retention_days": null,
      "is_announcement_only": false,
      "can_remove_subscribers_group": 2
    }
  ]
}`

// GetUsersResponse is a mock response for GET /api/v1/users.
const GetUsersResponse = `{
  "result": "success",
  "msg": "",
  "members": [
    {
      "user_id": 100,
      "delivery_email": null,
      "email": "user1@example.com",
      "full_name": "User One",
      "date_joined": "2024-01-01T00:00:00.000000Z",
      "is_active": true,
      "is_owner": false,
      "is_admin": false,
      "is_guest": false,
      "is_billing_admin": false,
      "is_bot": false,
      "bot_type": null,
      "timezone": "UTC",
      "avatar_url": "https://example.com/avatar/100.png",
      "avatar_version": 1,
      "profile_data": {}
    },
    {
      "user_id": 101,
      "delivery_email": null,
      "email": "user2@example.com",
      "full_name": "User Two",
      "date_joined": "2024-01-02T00:00:00.000000Z",
      "is_active": true,
      "is_owner": false,
      "is_admin": true,
      "is_guest": false,
      "is_billing_admin": false,
      "is_bot": false,
      "bot_type": null,
      "timezone": "America/New_York",
      "avatar_url": "https://example.com/avatar/101.png",
      "avatar_version": 1,
      "profile_data": {}
    },
    {
      "user_id": 102,
      "delivery_email": null,
      "email": "bot@example.com",
      "full_name": "Test Bot",
      "date_joined": "2024-01-03T00:00:00.000000Z",
      "is_active": true,
      "is_owner": false,
      "is_admin": false,
      "is_guest": false,
      "is_billing_admin": false,
      "is_bot": true,
      "bot_type": 1,
      "bot_owner_id": 101,
      "timezone": "UTC",
      "avatar_url": "https://example.com/avatar/102.png",
      "avatar_version": 1,
      "profile_data": {}
    }
  ]
}`

// DeleteQueueResponse is a mock response for DELETE /api/v1/events.
const DeleteQueueResponse = `{
  "result": "success",
  "msg": ""
}`
