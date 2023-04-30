import { defineStore } from 'pinia';
import { asyncFaillable } from '~/lib/utils';

type UserState = {
  id: string;
  username: string;
  token: string;
  status: 'logged-in' | 'logged-out';
};

const initialState: UserState = {
  id: '',
  username: '',
  token: '',
  status: 'logged-out',
};

export const useUserStore = defineStore('user', {
  state: (): UserState => initialState,
  actions: {
    async authenticate() {
      const { $httpClient } = useNuxtApp();
      const headers = useRequestHeaders(['cookie']);
      const response = await asyncFaillable(
        $httpClient.post<{ id: number; username: string; token: string }>({
          url: '/api/auth',
          options: { headers },
        }),
      );

      if (response.failed) {
        throw new Error('[store] Authentication failed');
      }

      this.$patch({
        id: response.result.id.toString(),
        username: response.result.username,
        token: response.result.token,
        status: 'logged-in',
      });
    },
    async register(credentials: { username: string; password: string }) {
      const { $httpClient } = useNuxtApp();
      const response = await asyncFaillable(
        $httpClient.post<{ id: number; username: string; token: string }>({
          url: '/api/register',
          data: {
            username: credentials.username,
            password: credentials.password,
          },
        }),
      );

      if (response.failed) {
        throw new Error('[store] Registration failed');
      }

      const config = useRuntimeConfig().public;
      const cookie = useCookie(config.cookieName);
      const state: UserState = {
        id: response.result.id.toString(),
        username: response.result.username,
        token: response.result.token,
        status: 'logged-in',
      };
      this.$patch(state);
      cookie.value = state.token;
    },
  },
});
