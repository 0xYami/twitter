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
      const config = useRuntimeConfig().public;
      const headers = useRequestHeaders(['cookie']);
      const response = await asyncFaillable<{ id: number; username: string; token: string }>(
        $fetch(`${config.serverBaseURL}/auth`, {
          method: 'POST',
          headers,
          parseResponse: JSON.parse,
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
      const config = useRuntimeConfig().public;
      const response = await asyncFaillable<{ id: number; username: string; token: string }>(
        $fetch(`${config.serverBaseURL}/register`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: {
            username: credentials.username,
            password: credentials.password,
          },
          parseResponse: JSON.parse,
        }),
      );

      if (response.failed) {
        throw new Error('[store] Registration failed');
      }

      const cookie = useCookie(config.cookieName, { httpOnly: true, secure: true });
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
