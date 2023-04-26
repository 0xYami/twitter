import { z } from 'zod';

const configSchema = z.object({
  cookieName: z.string().min(1),
  serverBaseURL: z.string().url(),
});

export const config = configSchema.parse({
  cookieName: process.env.COOKIE_NAME,
  serverBaseURL: process.env.SERVER_BASE_URL,
});
