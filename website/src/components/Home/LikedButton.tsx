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

interface Props {
  setLikedMangas: () => void;
}

export const LikedButton = ({ ...props }: Props) => (
  <button
    type="button"
    className="bg-yellow-300 hover:bg-yellow-400 text-gray-800 font-bold py-2 px-4 border-yellow-600 hover:border-yellow-500 inline-flex items-center w-56 mb-5"
    onClick={() => props.setLikedMangas()}
  >
    <svg
      xmlns="http://www.w3.org/2000/svg"
      className="h-5 w-5 cursor-pointer"
      fill="red"
      viewBox="0 0 24 24"
      stroke="red"
    >
      <path
        strokeLinecap="round"
        strokeLinejoin="round"
        strokeWidth={1}
        d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"
      />
    </svg>
    <span>Liked</span>
  </button>
);
