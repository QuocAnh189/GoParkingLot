import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

//interface
import {
  IExitResponse,
  IIOHistory,
  IListIOHistoryRequest,
} from "@interfaces/io_history";
import { IListData } from "@interfaces/common";

export const apiInOut = createApi({
  reducerPath: "apiInOut",
  baseQuery: fetchBaseQuery({
    baseUrl: import.meta.env.VITE_API_URL,
    prepareHeaders: (headers) => {
      const token = JSON.parse(localStorage.getItem("token")!)?.accessToken;

      if (token) {
        headers.set("Authorization", `Bearer ${token}`);
      }

      return headers;
    },
  }),
  keepUnusedDataFor: 20,
  tagTypes: ["InOut"],
  endpoints: (builder) => ({
    getIOHistories: builder.query<IListData<IIOHistory>, IListIOHistoryRequest>(
      {
        query: (params) => ({
          url: "/io-histories/",
          method: "GET",
          params,
        }),
        transformResponse: (response: any) => response.data,
        transformErrorResponse: (error) => error.data,
        providesTags: ["InOut"],
      }
    ),

    entrance: builder.mutation<IIOHistory, FormData>({
      query: (data) => ({
        url: `/io-histories/entrance`,
        method: "POST",
        body: data,
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      invalidatesTags: ["InOut"],
    }),

    exit: builder.mutation<IExitResponse, FormData>({
      query: (data) => ({
        url: `/io-histories/exit`,
        method: "POST",
        body: data,
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      invalidatesTags: ["InOut"],
    }),
  }),
});

export const { useGetIOHistoriesQuery, useEntranceMutation, useExitMutation } =
  apiInOut;
