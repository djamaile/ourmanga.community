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

import { useStore } from '../../global';

interface Props {
  image: string;
  name: string;
}

export const PublisherLogo = ({ ...props }: Props) => {
  const changePublisher = useStore(state => state.changePublisher);
  const publisher = useStore(state => state.publisher);
  const styles = { divClass: 'object-contain cursor-pointer m-auto block' };

  return (
    <section
      className="flex justify-center"
      onClick={() => changePublisher(props.name)}
    >
      <img
        src={props.image}
        alt={props.name}
        width="100"
        height="100"
        className={`${styles.divClass} 
          ${publisher !== props.name ? 'filter grayscale' : ''}`}
      />
    </section>
  );
};
