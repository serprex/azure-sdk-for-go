//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmediaservices

import "encoding/json"

func unmarshalAudioAnalyzerPresetClassification(rawMsg json.RawMessage) (AudioAnalyzerPresetClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AudioAnalyzerPresetClassification
	switch m["@odata.type"] {
	case "#Microsoft.Media.VideoAnalyzerPreset":
		b = &VideoAnalyzerPreset{}
	default:
		b = &AudioAnalyzerPreset{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalAudioAnalyzerPresetClassificationArray(rawMsg json.RawMessage) ([]AudioAnalyzerPresetClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]AudioAnalyzerPresetClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalAudioAnalyzerPresetClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalAudioClassification(rawMsg json.RawMessage) (AudioClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AudioClassification
	switch m["@odata.type"] {
	case "#Microsoft.Media.AacAudio":
		b = &AacAudio{}
	default:
		b = &Audio{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalAudioClassificationArray(rawMsg json.RawMessage) ([]AudioClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]AudioClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalAudioClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalAudioTrackDescriptorClassification(rawMsg json.RawMessage) (AudioTrackDescriptorClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AudioTrackDescriptorClassification
	switch m["@odata.type"] {
	case "#Microsoft.Media.SelectAudioTrackByAttribute":
		b = &SelectAudioTrackByAttribute{}
	case "#Microsoft.Media.SelectAudioTrackById":
		b = &SelectAudioTrackByID{}
	default:
		b = &AudioTrackDescriptor{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalAudioTrackDescriptorClassificationArray(rawMsg json.RawMessage) ([]AudioTrackDescriptorClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]AudioTrackDescriptorClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalAudioTrackDescriptorClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalClipTimeClassification(rawMsg json.RawMessage) (ClipTimeClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ClipTimeClassification
	switch m["@odata.type"] {
	case "#Microsoft.Media.AbsoluteClipTime":
		b = &AbsoluteClipTime{}
	case "#Microsoft.Media.UtcClipTime":
		b = &UTCClipTime{}
	default:
		b = &ClipTime{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalClipTimeClassificationArray(rawMsg json.RawMessage) ([]ClipTimeClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ClipTimeClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalClipTimeClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalCodecClassification(rawMsg json.RawMessage) (CodecClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b CodecClassification
	switch m["@odata.type"] {
	case "#Microsoft.Media.AacAudio":
		b = &AacAudio{}
	case "#Microsoft.Media.Audio":
		b = &Audio{}
	case "#Microsoft.Media.CopyAudio":
		b = &CopyAudio{}
	case "#Microsoft.Media.CopyVideo":
		b = &CopyVideo{}
	case "#Microsoft.Media.H264Video":
		b = &H264Video{}
	case "#Microsoft.Media.H265Video":
		b = &H265Video{}
	case "#Microsoft.Media.Image":
		b = &Image{}
	case "#Microsoft.Media.JpgImage":
		b = &JpgImage{}
	case "#Microsoft.Media.PngImage":
		b = &PNGImage{}
	case "#Microsoft.Media.Video":
		b = &Video{}
	default:
		b = &Codec{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalCodecClassificationArray(rawMsg json.RawMessage) ([]CodecClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]CodecClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalCodecClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalContentKeyPolicyConfigurationClassification(rawMsg json.RawMessage) (ContentKeyPolicyConfigurationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ContentKeyPolicyConfigurationClassification
	switch m["@odata.type"] {
	case "#Microsoft.Media.ContentKeyPolicyClearKeyConfiguration":
		b = &ContentKeyPolicyClearKeyConfiguration{}
	case "#Microsoft.Media.ContentKeyPolicyFairPlayConfiguration":
		b = &ContentKeyPolicyFairPlayConfiguration{}
	case "#Microsoft.Media.ContentKeyPolicyPlayReadyConfiguration":
		b = &ContentKeyPolicyPlayReadyConfiguration{}
	case "#Microsoft.Media.ContentKeyPolicyUnknownConfiguration":
		b = &ContentKeyPolicyUnknownConfiguration{}
	case "#Microsoft.Media.ContentKeyPolicyWidevineConfiguration":
		b = &ContentKeyPolicyWidevineConfiguration{}
	default:
		b = &ContentKeyPolicyConfiguration{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalContentKeyPolicyConfigurationClassificationArray(rawMsg json.RawMessage) ([]ContentKeyPolicyConfigurationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ContentKeyPolicyConfigurationClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalContentKeyPolicyConfigurationClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalContentKeyPolicyPlayReadyContentKeyLocationClassification(rawMsg json.RawMessage) (ContentKeyPolicyPlayReadyContentKeyLocationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ContentKeyPolicyPlayReadyContentKeyLocationClassification
	switch m["@odata.type"] {
	case "#Microsoft.Media.ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeader":
		b = &ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeader{}
	case "#Microsoft.Media.ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifier":
		b = &ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifier{}
	default:
		b = &ContentKeyPolicyPlayReadyContentKeyLocation{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalContentKeyPolicyPlayReadyContentKeyLocationClassificationArray(rawMsg json.RawMessage) ([]ContentKeyPolicyPlayReadyContentKeyLocationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ContentKeyPolicyPlayReadyContentKeyLocationClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalContentKeyPolicyPlayReadyContentKeyLocationClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalContentKeyPolicyRestrictionClassification(rawMsg json.RawMessage) (ContentKeyPolicyRestrictionClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ContentKeyPolicyRestrictionClassification
	switch m["@odata.type"] {
	case "#Microsoft.Media.ContentKeyPolicyOpenRestriction":
		b = &ContentKeyPolicyOpenRestriction{}
	case "#Microsoft.Media.ContentKeyPolicyTokenRestriction":
		b = &ContentKeyPolicyTokenRestriction{}
	case "#Microsoft.Media.ContentKeyPolicyUnknownRestriction":
		b = &ContentKeyPolicyUnknownRestriction{}
	default:
		b = &ContentKeyPolicyRestriction{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalContentKeyPolicyRestrictionClassificationArray(rawMsg json.RawMessage) ([]ContentKeyPolicyRestrictionClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ContentKeyPolicyRestrictionClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalContentKeyPolicyRestrictionClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalContentKeyPolicyRestrictionTokenKeyClassification(rawMsg json.RawMessage) (ContentKeyPolicyRestrictionTokenKeyClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ContentKeyPolicyRestrictionTokenKeyClassification
	switch m["@odata.type"] {
	case "#Microsoft.Media.ContentKeyPolicyRsaTokenKey":
		b = &ContentKeyPolicyRsaTokenKey{}
	case "#Microsoft.Media.ContentKeyPolicySymmetricTokenKey":
		b = &ContentKeyPolicySymmetricTokenKey{}
	case "#Microsoft.Media.ContentKeyPolicyX509CertificateTokenKey":
		b = &ContentKeyPolicyX509CertificateTokenKey{}
	default:
		b = &ContentKeyPolicyRestrictionTokenKey{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalContentKeyPolicyRestrictionTokenKeyClassificationArray(rawMsg json.RawMessage) ([]ContentKeyPolicyRestrictionTokenKeyClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ContentKeyPolicyRestrictionTokenKeyClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalContentKeyPolicyRestrictionTokenKeyClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalFormatClassification(rawMsg json.RawMessage) (FormatClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b FormatClassification
	switch m["@odata.type"] {
	case "#Microsoft.Media.ImageFormat":
		b = &ImageFormat{}
	case "#Microsoft.Media.JpgFormat":
		b = &JpgFormat{}
	case "#Microsoft.Media.Mp4Format":
		b = &Mp4Format{}
	case "#Microsoft.Media.MultiBitrateFormat":
		b = &MultiBitrateFormat{}
	case "#Microsoft.Media.PngFormat":
		b = &PNGFormat{}
	case "#Microsoft.Media.TransportStreamFormat":
		b = &TransportStreamFormat{}
	default:
		b = &Format{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalFormatClassificationArray(rawMsg json.RawMessage) ([]FormatClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]FormatClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalFormatClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalH265VideoLayerClassification(rawMsg json.RawMessage) (H265VideoLayerClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b H265VideoLayerClassification
	switch m["@odata.type"] {
	case "#Microsoft.Media.H265Layer":
		b = &H265Layer{}
	default:
		b = &H265VideoLayer{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalH265VideoLayerClassificationArray(rawMsg json.RawMessage) ([]H265VideoLayerClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]H265VideoLayerClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalH265VideoLayerClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalImageClassification(rawMsg json.RawMessage) (ImageClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ImageClassification
	switch m["@odata.type"] {
	case "#Microsoft.Media.JpgImage":
		b = &JpgImage{}
	case "#Microsoft.Media.PngImage":
		b = &PNGImage{}
	default:
		b = &Image{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalImageClassificationArray(rawMsg json.RawMessage) ([]ImageClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ImageClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalImageClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalImageFormatClassification(rawMsg json.RawMessage) (ImageFormatClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ImageFormatClassification
	switch m["@odata.type"] {
	case "#Microsoft.Media.JpgFormat":
		b = &JpgFormat{}
	case "#Microsoft.Media.PngFormat":
		b = &PNGFormat{}
	default:
		b = &ImageFormat{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalImageFormatClassificationArray(rawMsg json.RawMessage) ([]ImageFormatClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ImageFormatClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalImageFormatClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalInputDefinitionClassification(rawMsg json.RawMessage) (InputDefinitionClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b InputDefinitionClassification
	switch m["@odata.type"] {
	case "#Microsoft.Media.FromAllInputFile":
		b = &FromAllInputFile{}
	case "#Microsoft.Media.FromEachInputFile":
		b = &FromEachInputFile{}
	case "#Microsoft.Media.InputFile":
		b = &InputFile{}
	default:
		b = &InputDefinition{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalInputDefinitionClassificationArray(rawMsg json.RawMessage) ([]InputDefinitionClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]InputDefinitionClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalInputDefinitionClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalJobInputClassification(rawMsg json.RawMessage) (JobInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b JobInputClassification
	switch m["@odata.type"] {
	case "#Microsoft.Media.JobInputAsset":
		b = &JobInputAsset{}
	case "#Microsoft.Media.JobInputClip":
		b = &JobInputClip{}
	case "#Microsoft.Media.JobInputHttp":
		b = &JobInputHTTP{}
	case "#Microsoft.Media.JobInputSequence":
		b = &JobInputSequence{}
	case "#Microsoft.Media.JobInputs":
		b = &JobInputs{}
	default:
		b = &JobInput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalJobInputClassificationArray(rawMsg json.RawMessage) ([]JobInputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]JobInputClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalJobInputClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalJobInputClipClassification(rawMsg json.RawMessage) (JobInputClipClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b JobInputClipClassification
	switch m["@odata.type"] {
	case "#Microsoft.Media.JobInputAsset":
		b = &JobInputAsset{}
	case "#Microsoft.Media.JobInputHttp":
		b = &JobInputHTTP{}
	default:
		b = &JobInputClip{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalJobInputClipClassificationArray(rawMsg json.RawMessage) ([]JobInputClipClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]JobInputClipClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalJobInputClipClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalJobOutputClassification(rawMsg json.RawMessage) (JobOutputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b JobOutputClassification
	switch m["@odata.type"] {
	case "#Microsoft.Media.JobOutputAsset":
		b = &JobOutputAsset{}
	default:
		b = &JobOutput{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalJobOutputClassificationArray(rawMsg json.RawMessage) ([]JobOutputClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]JobOutputClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalJobOutputClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalLayerClassification(rawMsg json.RawMessage) (LayerClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b LayerClassification
	switch m["@odata.type"] {
	case "#Microsoft.Media.H264Layer":
		b = &H264Layer{}
	case "#Microsoft.Media.H265Layer":
		b = &H265Layer{}
	case "#Microsoft.Media.H265VideoLayer":
		b = &H265VideoLayer{}
	case "#Microsoft.Media.JpgLayer":
		b = &JpgLayer{}
	case "#Microsoft.Media.PngLayer":
		b = &PNGLayer{}
	case "#Microsoft.Media.VideoLayer":
		b = &VideoLayer{}
	default:
		b = &Layer{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalLayerClassificationArray(rawMsg json.RawMessage) ([]LayerClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]LayerClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalLayerClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalMultiBitrateFormatClassification(rawMsg json.RawMessage) (MultiBitrateFormatClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b MultiBitrateFormatClassification
	switch m["@odata.type"] {
	case "#Microsoft.Media.Mp4Format":
		b = &Mp4Format{}
	case "#Microsoft.Media.TransportStreamFormat":
		b = &TransportStreamFormat{}
	default:
		b = &MultiBitrateFormat{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalMultiBitrateFormatClassificationArray(rawMsg json.RawMessage) ([]MultiBitrateFormatClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]MultiBitrateFormatClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalMultiBitrateFormatClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalOverlayClassification(rawMsg json.RawMessage) (OverlayClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b OverlayClassification
	switch m["@odata.type"] {
	case "#Microsoft.Media.AudioOverlay":
		b = &AudioOverlay{}
	case "#Microsoft.Media.VideoOverlay":
		b = &VideoOverlay{}
	default:
		b = &Overlay{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalOverlayClassificationArray(rawMsg json.RawMessage) ([]OverlayClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]OverlayClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalOverlayClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalPresetClassification(rawMsg json.RawMessage) (PresetClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b PresetClassification
	switch m["@odata.type"] {
	case "#Microsoft.Media.AudioAnalyzerPreset":
		b = &AudioAnalyzerPreset{}
	case "#Microsoft.Media.BuiltInStandardEncoderPreset":
		b = &BuiltInStandardEncoderPreset{}
	case "#Microsoft.Media.FaceDetectorPreset":
		b = &FaceDetectorPreset{}
	case "#Microsoft.Media.StandardEncoderPreset":
		b = &StandardEncoderPreset{}
	case "#Microsoft.Media.VideoAnalyzerPreset":
		b = &VideoAnalyzerPreset{}
	default:
		b = &Preset{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalPresetClassificationArray(rawMsg json.RawMessage) ([]PresetClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]PresetClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalPresetClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalTrackDescriptorClassification(rawMsg json.RawMessage) (TrackDescriptorClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b TrackDescriptorClassification
	switch m["@odata.type"] {
	case "#Microsoft.Media.AudioTrackDescriptor":
		b = &AudioTrackDescriptor{}
	case "#Microsoft.Media.SelectAudioTrackByAttribute":
		b = &SelectAudioTrackByAttribute{}
	case "#Microsoft.Media.SelectAudioTrackById":
		b = &SelectAudioTrackByID{}
	case "#Microsoft.Media.SelectVideoTrackByAttribute":
		b = &SelectVideoTrackByAttribute{}
	case "#Microsoft.Media.SelectVideoTrackById":
		b = &SelectVideoTrackByID{}
	case "#Microsoft.Media.VideoTrackDescriptor":
		b = &VideoTrackDescriptor{}
	default:
		b = &TrackDescriptor{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalTrackDescriptorClassificationArray(rawMsg json.RawMessage) ([]TrackDescriptorClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]TrackDescriptorClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalTrackDescriptorClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalVideoClassification(rawMsg json.RawMessage) (VideoClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b VideoClassification
	switch m["@odata.type"] {
	case "#Microsoft.Media.H264Video":
		b = &H264Video{}
	case "#Microsoft.Media.H265Video":
		b = &H265Video{}
	case "#Microsoft.Media.Image":
		b = &Image{}
	case "#Microsoft.Media.JpgImage":
		b = &JpgImage{}
	case "#Microsoft.Media.PngImage":
		b = &PNGImage{}
	default:
		b = &Video{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalVideoClassificationArray(rawMsg json.RawMessage) ([]VideoClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]VideoClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalVideoClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalVideoLayerClassification(rawMsg json.RawMessage) (VideoLayerClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b VideoLayerClassification
	switch m["@odata.type"] {
	case "#Microsoft.Media.H264Layer":
		b = &H264Layer{}
	default:
		b = &VideoLayer{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalVideoLayerClassificationArray(rawMsg json.RawMessage) ([]VideoLayerClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]VideoLayerClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalVideoLayerClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalVideoTrackDescriptorClassification(rawMsg json.RawMessage) (VideoTrackDescriptorClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b VideoTrackDescriptorClassification
	switch m["@odata.type"] {
	case "#Microsoft.Media.SelectVideoTrackByAttribute":
		b = &SelectVideoTrackByAttribute{}
	case "#Microsoft.Media.SelectVideoTrackById":
		b = &SelectVideoTrackByID{}
	default:
		b = &VideoTrackDescriptor{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalVideoTrackDescriptorClassificationArray(rawMsg json.RawMessage) ([]VideoTrackDescriptorClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]VideoTrackDescriptorClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalVideoTrackDescriptorClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}
