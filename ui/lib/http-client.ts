import axios from 'axios';
import type { Axios, AxiosRequestConfig, CreateAxiosDefaults } from 'axios';
import { asyncFaillable } from './utils'

type GetConfig = {
  url: `/api/${string}`;
  options?: AxiosRequestConfig;
}

type PostConfig<TData> = {
  url: `/api/${string}`;
  data?: TData;
  options?: AxiosRequestConfig;
}

export class HttpClient {
  #http: Axios;

  constructor(config?: CreateAxiosDefaults) {
    this.#http = axios.create(config)
  }

  async get<T>(config: GetConfig): Promise<T> {
    const tryCall = await asyncFaillable<{ data: T }>(this.#http.get(config.url, config.options));
    if (tryCall.failed) {
      throw new Error(`GET: request to ${config.url} failed`)
    }
    return tryCall.result.data;
  }

  async post<TResponse, TData = unknown>(config: PostConfig<TData>): Promise<TResponse> {
    const tryCall = await asyncFaillable<{ data: TResponse }>(
      this.#http.post(config.url, config.data, config.options)
    );
    if (tryCall.failed) {
      throw new Error(`POST: request to ${config.url} failed`)
    }
    return tryCall.result.data;
  }
}
