package utils

import (
	"encoding/json"

	"github.com/rs/zerolog/log"

	"github.com/golang/protobuf/jsonpb"

	interaction "../../../proto/interaction"
)

// ConvertRawJSONToInteractionMessage TODO
func ConvertRawJSONToInteractionMessage(jsonRaw []byte) *interaction.InteractionMessage {
	dat := make(map[string]interface{})
	if err := json.Unmarshal(jsonRaw, &dat); err != nil {
		log.Error().Err(err).Msg("unable to ConvertRawJSONToInteractionMessage: Unmarshal jsonRaw")
	}

	jsonFrameRaw, err := json.Marshal(dat["frame"])
	if err != nil {
		log.Error().Err(err).Msg("unable to ConvertRawJSONToInteractionMessage: Marshal dat[\"frame\"]")
	}

	jsonFrame := string(jsonFrameRaw)

	protoFrame := &interaction.Frame{}
	jsonpb.UnmarshalString(jsonFrame, protoFrame)

	interactionElementsRaw, err := json.Marshal(dat["interactionElements"])
	if err != nil {
		log.Error().Err(err).Msg("unable to ConvertRawJSONToInteractionMessage: Marshal dat[\"interactionElements\"]")
	}

	protoMessage := &interaction.InteractionMessage{
		Frame:               protoFrame,
		InteractionElements: interactionElementsRaw,
	}

	return protoMessage
}

// ConvertInteractionMessageToRawJSON TODO
func ConvertInteractionMessageToRawJSON(protoMessage *interaction.InteractionMessage) []byte {
	protoFrame := protoMessage.Frame
	marshaler := jsonpb.Marshaler{}
	jsonFrame, err := marshaler.MarshalToString(protoFrame)
	if err != nil {
		log.Error().Err(err).Msg("unable to ConvertInteractionMessageToRawJSON: MarshalToString protoFrame")
	}

	interactionElementsRaw := protoMessage.InteractionElements

	jsonString := "{\"frame\":" + jsonFrame + ",\"interactionElements\":" + string(interactionElementsRaw) + "}"
	jsonRaw := []byte(jsonString)

	return jsonRaw
}
