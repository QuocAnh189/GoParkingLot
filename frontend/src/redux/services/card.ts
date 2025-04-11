import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

//interface
import {
  CreateCardPayload,
  ICard,
  IListCardRequest,
  UpdateCardPayload,
} from "@interfaces/card";
import { IListData } from "@interfaces/common";

export const apiCard = createApi({
  reducerPath: "apiCard",
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
  keepUnusedDataFor: 1,
  tagTypes: ["Card"],
  endpoints: (builder) => ({
    getCards: builder.query<IListData<ICard>, IListCardRequest>({
      query: (params) => ({
        url: "/cards/",
        method: "GET",
        params,
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      providesTags: ["Card"],
    }),

    getCard: builder.query<ICard, string>({
      query: (id) => ({
        url: `/cards/${id}`,
        method: "GET",
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      providesTags: ["Card"],
    }),

    createCard: builder.mutation<boolean, CreateCardPayload>({
      query: (data) => ({
        url: "/cards/",
        method: "POST",
        body: data,
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      invalidatesTags: ["Card"],
    }),

    updateCard: builder.mutation<boolean, UpdateCardPayload>({
      query: (data) => ({
        url: `/cards/${data.id}`,
        method: "PUT",
        body: data,
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      invalidatesTags: ["Card"],
    }),

    deleteCard: builder.mutation<boolean, string>({
      query: (id) => ({
        url: `/cards/${id}`,
        method: "DELETE",
      }),
      transformResponse: (response: any) => response.data,
      transformErrorResponse: (error) => error.data,
      invalidatesTags: ["Card"],
    }),
  }),
});

export const {
  useGetCardsQuery,
  useGetCardQuery,
  useCreateCardMutation,
  useUpdateCardMutation,
  useDeleteCardMutation,
} = apiCard;
