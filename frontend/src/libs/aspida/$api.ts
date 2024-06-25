import type { AspidaClient, BasicHeaders } from 'aspida';
import { dataToURLString } from 'aspida';
import type { Methods as Methods_by08hd } from '.';
import type { Methods as Methods_1r1gcz6 } from './hello';
import type { Methods as Methods_sjgy3w } from './tweets';
import type { Methods as Methods_ogekxv } from './tweets/_tweetId@number';

const api = <T>({ baseURL, fetch }: AspidaClient<T>) => {
  const prefix = (baseURL === undefined ? 'http://localhost:9000' : baseURL).replace(/\/$/, '');
  const PATH0 = '/hello';
  const PATH1 = '/tweets';
  const GET = 'GET';

  return {
    hello: {
      /**
       * Return Hello message
       * @returns A successful response
       */
      get: (option?: { query?: Methods_1r1gcz6['get']['query'] | undefined, config?: T | undefined } | undefined) =>
        fetch<Methods_1r1gcz6['get']['resBody'], BasicHeaders, Methods_1r1gcz6['get']['status']>(prefix, PATH0, GET, option).json(),
      /**
       * Return Hello message
       * @returns A successful response
       */
      $get: (option?: { query?: Methods_1r1gcz6['get']['query'] | undefined, config?: T | undefined } | undefined) =>
        fetch<Methods_1r1gcz6['get']['resBody'], BasicHeaders, Methods_1r1gcz6['get']['status']>(prefix, PATH0, GET, option).json().then(r => r.body),
      $path: (option?: { method?: 'get' | undefined; query: Methods_1r1gcz6['get']['query'] } | undefined) =>
        `${prefix}${PATH0}${option && option.query ? `?${dataToURLString(option.query)}` : ''}`,
    },
    tweets: {
      _tweetId: (val1: number) => {
        const prefix1 = `${PATH1}/${val1}`;

        return {
          /**
           * Return tweet
           * @returns A successful response
           */
          get: (option?: { config?: T | undefined } | undefined) =>
            fetch<Methods_ogekxv['get']['resBody'], BasicHeaders, Methods_ogekxv['get']['status']>(prefix, prefix1, GET, option).json(),
          /**
           * Return tweet
           * @returns A successful response
           */
          $get: (option?: { config?: T | undefined } | undefined) =>
            fetch<Methods_ogekxv['get']['resBody'], BasicHeaders, Methods_ogekxv['get']['status']>(prefix, prefix1, GET, option).json().then(r => r.body),
          $path: () => `${prefix}${prefix1}`,
        };
      },
      /**
       * Return tweets
       * @returns A successful response
       */
      get: (option?: { query?: Methods_sjgy3w['get']['query'] | undefined, config?: T | undefined } | undefined) =>
        fetch<Methods_sjgy3w['get']['resBody'], BasicHeaders, Methods_sjgy3w['get']['status']>(prefix, PATH1, GET, option).json(),
      /**
       * Return tweets
       * @returns A successful response
       */
      $get: (option?: { query?: Methods_sjgy3w['get']['query'] | undefined, config?: T | undefined } | undefined) =>
        fetch<Methods_sjgy3w['get']['resBody'], BasicHeaders, Methods_sjgy3w['get']['status']>(prefix, PATH1, GET, option).json().then(r => r.body),
      $path: (option?: { method?: 'get' | undefined; query: Methods_sjgy3w['get']['query'] } | undefined) =>
        `${prefix}${PATH1}${option && option.query ? `?${dataToURLString(option.query)}` : ''}`,
    },
    /**
     * Return root message
     * @returns A successful response
     */
    get: (option?: { config?: T | undefined } | undefined) =>
      fetch<Methods_by08hd['get']['resBody'], BasicHeaders, Methods_by08hd['get']['status']>(prefix, '', GET, option).text(),
    /**
     * Return root message
     * @returns A successful response
     */
    $get: (option?: { config?: T | undefined } | undefined) =>
      fetch<Methods_by08hd['get']['resBody'], BasicHeaders, Methods_by08hd['get']['status']>(prefix, '', GET, option).text().then(r => r.body),
    $path: () => `${prefix}`,
  };
};

export type ApiInstance = ReturnType<typeof api>;
export default api;