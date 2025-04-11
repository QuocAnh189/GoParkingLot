import { configureStore } from "@reduxjs/toolkit";

//service
import { apiAuth } from "./services/auth";
import { apiCard } from "./services/card";
import { apiInOut } from "./services/in_out";

//slice
import authReducer, { AuthSliceKey } from "./slices/auth";

const store = configureStore({
  reducer: {
    [apiAuth.reducerPath]: apiAuth.reducer,
    [apiCard.reducerPath]: apiCard.reducer,
    [apiInOut.reducerPath]: apiInOut.reducer,

    [AuthSliceKey]: authReducer,
  },

  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: false,
      immutableCheck: false,
    }).concat([apiAuth.middleware, apiCard.middleware, apiInOut.middleware]),
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

export default store;
