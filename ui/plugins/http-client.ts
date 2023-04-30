import { HttpClient } from '~/lib/http-client';

export default defineNuxtPlugin(() => {
  const config = useRuntimeConfig().public;
  const httpClient = new HttpClient({
    baseURL: config.serverBaseURL,
  });
  return {
    provide: {
      httpClient,
    },
  };
});
