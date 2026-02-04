import { GoogleGenAI, GenerateContentResponse } from "@google/genai";
import type { ChatMessage } from "@/types";

const apiKey = import.meta.env.VITE_API_KEY;
let ai: GoogleGenAI | null = null;

const getClient = () => {
    if (!apiKey) {
        return null;
    }
    if (!ai) {
        ai = new GoogleGenAI({ apiKey });
    }
    return ai;
};

const MODEL_NAME = "gemini-2.5-flash";

export const generateSummary = async (content: string): Promise<string> => {
    try {
        const client = getClient();
        if (!client) {
            return "AI disabled: missing VITE_API_KEY.";
        }
        const prompt = `Summarize the following blog post content in 2-3 concise sentences. Capture the main essence. \n\nContent: ${content.substring(0, 5000)}`; // Truncate to avoid huge context if necessary
        const response = await client.models.generateContent({
            model: MODEL_NAME,
            contents: prompt,
        });
        return response.text || "Could not generate summary.";
    } catch (error) {
        console.error("Gemini Summary Error:", error);
        return "Unable to generate summary at this time.";
    }
};

export const generateInsight = async (content: string): Promise<string> => {
    try {
        const client = getClient();
        if (!client) {
            return "AI disabled: missing VITE_API_KEY.";
        }
        const prompt = `Provide a unique, thought-provoking insight or a counter-point related to this blog post. Keep it brief (under 50 words) and engaging. \n\nContent: ${content.substring(0, 5000)}`;
        const response = await client.models.generateContent({
            model: MODEL_NAME,
            contents: prompt,
        });
        return response.text || "No insight available.";
    } catch (error) {
        console.error("Gemini Insight Error:", error);
        return "Unable to generate insight.";
    }
}

export const streamChatResponse = async (
    history: ChatMessage[],
    currentMessage: string,
    blogContext: string,
    onChunk: (text: string) => void
): Promise<void> => {
    try {
        const client = getClient();
        if (!client) {
            onChunk("AI disabled: missing VITE_API_KEY.");
            return;
        }
        const systemInstruction = `You are Lumina, an intelligent AI assistant for a personal blog. 
        You have access to the context of the blog posts the user is reading or asking about.
        
        Current Blog Context (Titles/Excerpts):
        ${blogContext}

        Your Tone:
        - Helpful, intellectual, yet warm and conversational.
        - Concise but insightful.
        - If asked about the author (Alex Chen), mention he is a creative developer and writer.
        
        Do not answer questions unrelated to the blog topics, technology, lifestyle, or general knowledge.
        `;

        const chat = client.chats.create({
            model: MODEL_NAME,
            config: {
                systemInstruction,
            },
            history: history.map(h => ({
                role: h.role,
                parts: [{ text: h.text }]
            }))
        });

        const result = await chat.sendMessageStream({ message: currentMessage });

        for await (const chunk of result) {
            const c = chunk as GenerateContentResponse;
            const text = c.text;
            if (text) {
                onChunk(text);
            }
        }
    } catch (error) {
        console.error("Chat Stream Error:", error);
        onChunk("\n[Connection Error: Unable to reach Gemini]");
    }
};