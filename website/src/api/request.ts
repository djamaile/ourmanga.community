// Copyright 2021 Djamaile Rahamat
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import axios, { Method, AxiosResponse } from 'axios';

const api = axios.create({
  baseURL: process.env.REACT_APP_HOST_BACKEND,
});

const request = <Type>(
  method: Method,
  url: string,
  params: Type,
): Promise<AxiosResponse<Type>> =>
  api.request<Type>({
    method,
    url,
    params,
  });

// Define a default query function that will receive the query key
// TODO(drahamat): figure out how to get rid of any type
export const defaultQueryFn = async <Type>({
  queryKey,
}: any): Promise<AxiosResponse<Type>> => {
  const data = await request(queryKey[0], queryKey[1], queryKey[2]);
  return data as AxiosResponse<Type>;
};
