package imitate

import (
	"strings"
)

func ConvertToString(chatgpt_response *ChatGPTResponse, previous_text *StringStruct, role bool) string {
	translated_response := NewChatCompletionChunk(strings.Replace(chatgpt_response.Message.Content.Parts[0].(string), previous_text.Text, "", 1))
	if role {
		translated_response.Choices[0].Delta.Role = chatgpt_response.Message.Author.Role
	} else if translated_response.Choices[0].Delta.Content == "" || (strings.HasPrefix(chatgpt_response.Message.Metadata.ModelSlug, "gpt-4") && translated_response.Choices[0].Delta.Content == "【") {
		return translated_response.Choices[0].Delta.Content
	}
	previous_text.Text = chatgpt_response.Message.Content.Parts[0].(string)
	return "data: " + translated_response.String() + "\n\n"
}