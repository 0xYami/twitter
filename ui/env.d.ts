declare namespace NodeJS {
  // eslint-disable-next-line @typescript-eslint/consistent-type-definitions
  interface ProcessEnv {
    NODE_ENV?: 'development' | 'production' | 'test' | undefined;
    SERVER_BASE_URL?: string | undefined;
  }
}
