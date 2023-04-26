import { z } from 'zod';

const configSchema = z.object({
  serverBaseURL: z.string().url(),
});

export const config = configSchema.parse({
  serverBaseURL: process.env.SERVER_BASE_URL,
});
