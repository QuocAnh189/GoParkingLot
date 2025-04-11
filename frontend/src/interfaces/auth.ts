export interface IUser {
  id: string;
  email: string;
  name: string;
  avatar_url: string;
}

export interface IAuth {
  accessToken: string;
  refreshToken: string;
  user: IUser;
}

export interface SingInPayload {
  email: string;
  password: string;
}

export interface SignUpPayload {
  email: string;
  name: string;
  avatar: any;
  password: string;
}
