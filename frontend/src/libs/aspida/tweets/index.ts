/* eslint-disable */
import type * as Types from '../@types'

export type Methods = {
  /** Return tweets */
  get: {
    status: 200
    /** A successful response */
    resBody: Types.TweetList
  }

  /** Return id and created time */
  post: {
    status: 201

    /** A successful response */
    resBody: {
      created_at?: string | undefined
    }

    reqBody: {
      title: string
      body: string
    }
  }
}
