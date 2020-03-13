/*
Copyright 2014 The go-marathon Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package marathon

import "fmt"

// EventType is a wrapper for a marathon event
type EventType struct {
	EventType string `json:"eventType"`
}

const (
	// EventIDAPIRequest is the event listener ID for the corresponding event.
	EventIDAPIRequest = 1 << iota
	// EventIDStatusUpdate is the event listener ID for the corresponding event.
	EventIDStatusUpdate
	// EventIDFrameworkMessage is the event listener ID for the corresponding event.
	EventIDFrameworkMessage
	// EventIDSubscription is the event listener ID for the corresponding event.
	EventIDSubscription
	// EventIDUnsubscribed is the event listener ID for the corresponding event.
	EventIDUnsubscribed
	// EventIDStreamAttached is the event listener ID for the corresponding event.
	EventIDStreamAttached
	// EventIDStreamDetached is the event listener ID for the corresponding event.
	EventIDStreamDetached
	// EventIDAddHealthCheck is the event listener ID for the corresponding event.
	EventIDAddHealthCheck
	// EventIDRemoveHealthCheck is the event listener ID for the corresponding event.
	EventIDRemoveHealthCheck
	// EventIDFailedHealthCheck is the event listener ID for the corresponding event.
	EventIDFailedHealthCheck
	// EventIDChangedHealthCheck is the event listener ID for the corresponding event.
	EventIDChangedHealthCheck
	// EventIDGroupChangeSuccess is the event listener ID for the corresponding event.
	EventIDGroupChangeSuccess
	// EventIDGroupChangeFailed is the event listener ID for the corresponding event.
	EventIDGroupChangeFailed
	// EventIDDeploymentSuccess is the event listener ID for the corresponding event.
	EventIDDeploymentSuccess
	// EventIDDeploymentFailed is the event listener ID for the corresponding event.
	EventIDDeploymentFailed
	// EventIDDeploymentInfo is the event listener ID for the corresponding event.
	EventIDDeploymentInfo
	// EventIDDeploymentStepSuccess is the event listener ID for the corresponding event.
	EventIDDeploymentStepSuccess
	// EventIDDeploymentStepFailed is the event listener ID for the corresponding event.
	EventIDDeploymentStepFailed
	// EventIDAppTerminated is the event listener ID for the corresponding event.
	EventIDAppTerminated
	// EventIDPodCreated is the event listener ID for the corresponding event.
	EventIDPodCreated
	// EventIDPodUpdated is the event listener ID for the corresponding event.
	EventIDPodUpdated
	// EventIDPodDeleted is the event listener ID for the corresponding event.
	EventIDPodDeleted
	// EventIDUnhealthyTaskKill is the event listener ID for the corresponding event.
	EventIDUnhealthyTaskKill
	// EventIDUnhealthyInstanceKill is the event listener ID for the corresponding event.
	EventIDUnhealthyInstanceKill
	// EventIDInstanceChanged is the event listener ID for the corresponding event.
	EventIDInstanceChanged
	// EventIDUnknownInstanceTerminated is the event listener ID for the corresponding event.
	EventIDUnknownInstanceTerminated
	// EventIDInstanceHealthChanged is the event listener ID for the corresponding event.
	EventIDInstanceHealthChanged

	// EventIDOurStreamAttached fires when our own SSE cstream to marathon is attached. Not a marathon event
	EventIDOurStreamAttached
	// EventIDOurStreamDetached fires when our own SSE stream to marathon is detached. Not a marathon event
	EventIDOurStreamDetached

	//EventIDApplications comprises all listener IDs for application events.
	EventIDApplications = EventIDStatusUpdate | EventIDChangedHealthCheck | EventIDFailedHealthCheck | EventIDAppTerminated
	//EventIDSubscriptions comprises all listener IDs for subscription events.
	EventIDSubscriptions = EventIDSubscription | EventIDUnsubscribed | EventIDStreamAttached | EventIDStreamDetached
)

var (
	eventTypesMap map[string]int
)

func init() {
	eventTypesMap = map[string]int{
		"api_post_event":                    EventIDAPIRequest,
		"status_update_event":               EventIDStatusUpdate,
		"framework_message_event":           EventIDFrameworkMessage,
		"subscribe_event":                   EventIDSubscription,
		"unsubscribe_event":                 EventIDUnsubscribed,
		"event_stream_attached":             EventIDStreamAttached,
		"event_stream_detached":             EventIDStreamDetached,
		"add_health_check_event":            EventIDAddHealthCheck,
		"remove_health_check_event":         EventIDRemoveHealthCheck,
		"failed_health_check_event":         EventIDFailedHealthCheck,
		"health_status_changed_event":       EventIDChangedHealthCheck,
		"group_change_success":              EventIDGroupChangeSuccess,
		"group_change_failed":               EventIDGroupChangeFailed,
		"deployment_success":                EventIDDeploymentSuccess,
		"deployment_failed":                 EventIDDeploymentFailed,
		"deployment_info":                   EventIDDeploymentInfo,
		"deployment_step_success":           EventIDDeploymentStepSuccess,
		"deployment_step_failure":           EventIDDeploymentStepFailed,
		"app_terminated_event":              EventIDAppTerminated,
		"pod_created_event":                 EventIDPodCreated,
		"pod_updated_event":                 EventIDPodUpdated,
		"pod_deleted_event":                 EventIDPodDeleted,
		"unhealthy_task_kill_event":         EventIDUnhealthyTaskKill,
		"unhealthy_instance_kill_event":     EventIDUnhealthyInstanceKill,
		"instance_changed_event":            EventIDInstanceChanged,
		"unknown_instance_terminated_event": EventIDUnknownInstanceTerminated,
		"instance_health_changed_event":     EventIDInstanceHealthChanged,
		"our_stream_attached":               EventIDOurStreamAttached, // this is our own event, not marathons
		"our_stream_detached":               EventIDOurStreamDetached, // this is our own event, not marathons
	}
}

//
//  Events taken from: https://mesosphere.github.io/marathon/docs/event-bus.html
//

// Event is the definition for a event in marathon
type Event struct {
	ID    int
	Name  string
	Event interface{}
}

func (event *Event) String() string {
	return fmt.Sprintf("type: %s, event: %s", event.Name, event.Event)
}

// EventsChannel is a channel to receive events upon
type EventsChannel chan *Event

/* --- Our connection to marathon SSE --- */

// EventOurStreamAttached occurs when our SSE stream attaches
type EventOurStreamAttached struct {
	EventType string `json:"eventType"`
}

// EventOurStreamDetached occurs when our SSE stream detaches
type EventOurStreamDetached struct {
	EventType string `json:"eventType"`
}

/* --- API Request --- */

// EventAPIRequest describes an 'api_post_event' event.
type EventAPIRequest struct {
	EventType     string       `json:"eventType"`
	ClientIP      string       `json:"clientIp"`
	Timestamp     string       `json:"timestamp"`
	URI           string       `json:"uri"`
	AppDefinition *Application `json:"appDefinition"`
}

/* --- Status Update --- */

// EventStatusUpdate describes a 'status_update_event' event.
type EventStatusUpdate struct {
	EventType   string       `json:"eventType"`
	Timestamp   string       `json:"timestamp,omitempty"`
	SlaveID     string       `json:"slaveId,omitempty"`
	TaskID      string       `json:"taskId"`
	TaskStatus  string       `json:"taskStatus"`
	Message     string       `json:"message,omitempty"`
	AppID       string       `json:"appId"`
	Host        string       `json:"host"`
	Ports       []int        `json:"ports,omitempty"`
	IPAddresses []*IPAddress `json:"ipAddresses"`
	Version     string       `json:"version,omitempty"`
}

// EventAppTerminated describes an 'app_terminated_event' event.
type EventAppTerminated struct {
	EventType string `json:"eventType"`
	Timestamp string `json:"timestamp,omitempty"`
	AppID     string `json:"appId"`
}

// EventInstanceChanged describes a 'instance_changed_event' event
type EventInstanceChanged struct {
	EventType      string `json:"eventType"`
	Timestamp      string `json:"timestamp"`
	InstanceID     string `json:"instanceId"`
	Condition      string `json:"condition"`
	RunSpecID      string `json:"runSpecId"`
	AgentID        string `json:"agentId"`
	Host           string `json:"host"`
	RunSpecVersion string `json:"runSpecVersion"`
}

// EventUnknownInstanceTerminated describes a 'unknown_instance_terminated_event' event
type EventUnknownInstanceTerminated struct {
	EventType  string `json:"eventType"`
	Timestamp  string `json:"timestamp"`
	InstanceID string `json:"instanceId"`
	Condition  string `json:"condition"`
	RunSpecID  string `json:"runSpecId"`
}

// EventInstanceHealthChanged describes a 'instance_health_changed_event' event
type EventInstanceHealthChanged struct {
	EventType      string `json:"eventType"`
	Timestamp      string `json:"timestamp"`
	InstanceID     string `json:"instanceId"`
	RunSpecVersion string `json:"runSpecVersion"`
	Healthy        bool   `json:"healthy"`
	RunSpecID      string `json:"runSpecId"`
}

/* --- Framework Message --- */

// EventFrameworkMessage describes a 'framework_message_event' event.
type EventFrameworkMessage struct {
	EventType  string `json:"eventType"`
	ExecutorID string `json:"executorId"`
	Message    string `json:"message"`
	SlaveID    string `json:"slaveId"`
	Timestamp  string `json:"timestamp"`
}

/* --- Event Subscription --- */

// EventSubscription describes a 'subscribe_event' event.
type EventSubscription struct {
	CallbackURL string `json:"callbackUrl"`
	ClientIP    string `json:"clientIp"`
	EventType   string `json:"eventType"`
	Timestamp   string `json:"timestamp"`
}

// EventUnsubscription describes an 'unsubscribe_event' event.
type EventUnsubscription struct {
	CallbackURL string `json:"callbackUrl"`
	ClientIP    string `json:"clientIp"`
	EventType   string `json:"eventType"`
	Timestamp   string `json:"timestamp"`
}

// EventStreamAttached describes an 'event_stream_attached' event.
type EventStreamAttached struct {
	RemoteAddress string `json:"remoteAddress"`
	EventType     string `json:"eventType"`
	Timestamp     string `json:"timestamp"`
}

// EventStreamDetached describes an 'event_stream_detached' event.
type EventStreamDetached struct {
	RemoteAddress string `json:"remoteAddress"`
	EventType     string `json:"eventType"`
	Timestamp     string `json:"timestamp"`
}

/* --- Health Checks --- */

// EventAddHealthCheck describes an 'add_health_check_event' event.
type EventAddHealthCheck struct {
	AppID       string `json:"appId"`
	EventType   string `json:"eventType"`
	HealthCheck struct {
		GracePeriodSeconds     float64 `json:"gracePeriodSeconds"`
		IntervalSeconds        float64 `json:"intervalSeconds"`
		MaxConsecutiveFailures float64 `json:"maxConsecutiveFailures"`
		Path                   string  `json:"path"`
		PortIndex              float64 `json:"portIndex"`
		Protocol               string  `json:"protocol"`
		TimeoutSeconds         float64 `json:"timeoutSeconds"`
	} `json:"healthCheck"`
	Timestamp string `json:"timestamp"`
}

// EventRemoveHealthCheck describes a 'remove_health_check_event' event.
type EventRemoveHealthCheck struct {
	AppID       string `json:"appId"`
	EventType   string `json:"eventType"`
	HealthCheck struct {
		GracePeriodSeconds     float64 `json:"gracePeriodSeconds"`
		IntervalSeconds        float64 `json:"intervalSeconds"`
		MaxConsecutiveFailures float64 `json:"maxConsecutiveFailures"`
		Path                   string  `json:"path"`
		PortIndex              float64 `json:"portIndex"`
		Protocol               string  `json:"protocol"`
		TimeoutSeconds         float64 `json:"timeoutSeconds"`
	} `json:"healthCheck"`
	Timestamp string `json:"timestamp"`
}

// EventFailedHealthCheck describes a 'failed_health_check_event' event.
type EventFailedHealthCheck struct {
	AppID       string `json:"appId"`
	EventType   string `json:"eventType"`
	HealthCheck struct {
		GracePeriodSeconds     float64 `json:"gracePeriodSeconds"`
		IntervalSeconds        float64 `json:"intervalSeconds"`
		MaxConsecutiveFailures float64 `json:"maxConsecutiveFailures"`
		Path                   string  `json:"path"`
		PortIndex              float64 `json:"portIndex"`
		Protocol               string  `json:"protocol"`
		TimeoutSeconds         float64 `json:"timeoutSeconds"`
	} `json:"healthCheck"`
	Timestamp string `json:"timestamp"`
}

// EventHealthCheckChanged describes a 'health_status_changed_event' event.
type EventHealthCheckChanged struct {
	EventType  string `json:"eventType"`
	Timestamp  string `json:"timestamp,omitempty"`
	AppID      string `json:"appId,omitEmpty"`
	TaskID     string `json:"taskId,omitempty"`
	InstanceID string `json:"instanceId,omitempty"`
	Version    string `json:"version,omitempty"`
	Alive      bool   `json:"alive"`
}

// EventUnhealthyTaskKill describes a 'unhealthy_task_kill_event' event
type EventUnhealthyTaskKill struct {
	EventType string `json:"eventType"`
	Timestamp string `json:"timestamp"`
	AppID     string `json:"appId"`
	TaskID    string `json:"taskId"`
	Version   string `json:"version"`
	Reason    string `json:"reason"`
	Host      string `json:"host"`
	SlaveID   string `json:"slaveId"`
}

// EventUnhealthyInstanceKill describes a 'unhealthy_instance_kill_event' event
type EventUnhealthyInstanceKill struct {
	EventType  string `json:"eventType"`
	Timestamp  string `json:"timestamp"`
	AppID      string `json:"appId"`
	TaskID     string `json:"taskId"`
	InstanceID string `json:"instanceId"`
	Version    string `json:"version"`
	Reason     string `json:"reason"`
	Host       string `json:"host"`
	SlaveID    string `json:"slaveId"`
}

/* --- Deployments --- */

// EventGroupChangeSuccess describes a 'group_change_success' event.
type EventGroupChangeSuccess struct {
	EventType string `json:"eventType"`
	GroupID   string `json:"groupId"`
	Timestamp string `json:"timestamp"`
	Version   string `json:"version"`
}

// EventGroupChangeFailed describes a 'group_change_failed' event.
type EventGroupChangeFailed struct {
	EventType string `json:"eventType"`
	GroupID   string `json:"groupId"`
	Timestamp string `json:"timestamp"`
	Version   string `json:"version"`
	Reason    string `json:"reason"`
}

// EventDeploymentSuccess describes a 'deployment_success' event.
type EventDeploymentSuccess struct {
	ID        string          `json:"id"`
	EventType string          `json:"eventType"`
	Timestamp string          `json:"timestamp"`
	Plan      *DeploymentPlan `json:"plan"`
}

// EventDeploymentFailed describes a 'deployment_failed' event.
type EventDeploymentFailed struct {
	ID        string `json:"id"`
	EventType string `json:"eventType"`
	Timestamp string `json:"timestamp"`
}

// EventDeploymentInfo describes a 'deployment_info' event.
type EventDeploymentInfo struct {
	EventType   string          `json:"eventType"`
	CurrentStep *StepActions    `json:"currentStep"`
	Timestamp   string          `json:"timestamp"`
	Plan        *DeploymentPlan `json:"plan"`
}

// EventDeploymentStepSuccess describes a 'deployment_step_success' event.
type EventDeploymentStepSuccess struct {
	EventType   string          `json:"eventType"`
	CurrentStep *StepActions    `json:"currentStep"`
	Timestamp   string          `json:"timestamp"`
	Plan        *DeploymentPlan `json:"plan"`
}

// EventDeploymentStepFailure describes a 'deployment_step_failure' event.
type EventDeploymentStepFailure struct {
	EventType   string          `json:"eventType"`
	CurrentStep *StepActions    `json:"currentStep"`
	Timestamp   string          `json:"timestamp"`
	Plan        *DeploymentPlan `json:"plan"`
}

/* --- Pods --- */

// EventPodCreated describes a 'pod_created_event' event
type EventPodCreated struct {
	EventType string `json:"eventType"`
	Timestamp string `json:"timestamp"`
	ClientIP  string `json:"clientIp"`
	URI       string `json:"uri"`
}

// EventPodUpdated describes a 'pod_created_event' event
type EventPodUpdated struct {
	EventType string `json:"eventType"`
	Timestamp string `json:"timestamp"`
	ClientIP  string `json:"clientIp"`
	URI       string `json:"uri"`
}

// EventPodDeleted describes a 'pod_created_event' event
type EventPodDeleted struct {
	EventType string `json:"eventType"`
	Timestamp string `json:"timestamp"`
	ClientIP  string `json:"clientIp"`
	URI       string `json:"uri"`
}

// GetEvent returns allocated empty event object which corresponds to provided event type
//		eventType:			the type of Marathon event
func GetEvent(eventType string) (*Event, error) {
	// step: check it's supported
	id, found := eventTypesMap[eventType]
	if found {
		event := new(Event)
		event.ID = id
		event.Name = eventType
		switch eventType {
		case "api_post_event":
			event.Event = new(EventAPIRequest)
		case "status_update_event":
			event.Event = new(EventStatusUpdate)
		case "framework_message_event":
			event.Event = new(EventFrameworkMessage)
		case "subscribe_event":
			event.Event = new(EventSubscription)
		case "unsubscribe_event":
			event.Event = new(EventUnsubscription)
		case "event_stream_attached":
			event.Event = new(EventStreamAttached)
		case "event_stream_detached":
			event.Event = new(EventStreamDetached)
		case "add_health_check_event":
			event.Event = new(EventAddHealthCheck)
		case "remove_health_check_event":
			event.Event = new(EventRemoveHealthCheck)
		case "failed_health_check_event":
			event.Event = new(EventFailedHealthCheck)
		case "health_status_changed_event":
			event.Event = new(EventHealthCheckChanged)
		case "group_change_success":
			event.Event = new(EventGroupChangeSuccess)
		case "group_change_failed":
			event.Event = new(EventGroupChangeFailed)
		case "deployment_success":
			event.Event = new(EventDeploymentSuccess)
		case "deployment_failed":
			event.Event = new(EventDeploymentFailed)
		case "deployment_info":
			event.Event = new(EventDeploymentInfo)
		case "deployment_step_success":
			event.Event = new(EventDeploymentStepSuccess)
		case "deployment_step_failure":
			event.Event = new(EventDeploymentStepFailure)
		case "app_terminated_event":
			event.Event = new(EventAppTerminated)
		case "pod_created_event":
			event.Event = new(EventPodCreated)
		case "pod_updated_event":
			event.Event = new(EventPodUpdated)
		case "pod_deleted_event":
			event.Event = new(EventPodDeleted)
		case "unhealthy_task_kill_event":
			event.Event = new(EventUnhealthyTaskKill)
		case "unhealthy_instance_kill_event":
			event.Event = new(EventUnhealthyInstanceKill)
		case "instance_changed_event":
			event.Event = new(EventInstanceChanged)
		case "unknown_instance_terminated_event":
			event.Event = new(EventUnknownInstanceTerminated)
		case "instance_health_changed_event":
			event.Event = new(EventInstanceHealthChanged)
		case "our_stream_attached":
			event.Event = new(EventOurStreamAttached)
		case "our_stream_detached":
			event.Event = new(EventOurStreamDetached)
		}

		return event, nil
	}

	return nil, fmt.Errorf("the event type: %s was not found or supported", eventType)
}
