import { useUserStore } from '~/stores/user';
import { asyncFaillable } from '~/lib/utils';

export default defineNuxtRouteMiddleware(async (to) => {
  const config = useRuntimeConfig().public;
  const cookie = useCookie(config.cookieName);
  if (!cookie.value && to.path !== '/register') {
    return navigateTo('register');
  }

  const user = useUserStore();

  if (user.status === 'logged-in') return;
  // TODO: properly check if user is logged in
  if (to.path === '/register') return;

  const tryAuth = await asyncFaillable(user.authenticate());
  if (tryAuth.failed) {
    return navigateTo('register');
  }
});
