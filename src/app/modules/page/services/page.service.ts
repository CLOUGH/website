import { Injectable } from '@angular/core';
import { of, Observable, range } from 'rxjs';
import * as faker from 'faker';

import { Post } from '../../post/post';
import { ILink } from '../../../store/link/link';

@Injectable({
  providedIn: 'root'
})
export class PageService {
  pages = [{
    path: '',
    name: 'Home',
    sections: [
      {
        hero: {
          image: 'assets/marius-masalar-132751-unsplash.jpg',
          mainHeading: 'Hi Warren Clough',
          subHeading: 'loremUt consectetur consequat non pariatur.',
          links: [
            {
              text: 'Contact Me',
              link: '#'
            }
          ]
        }
      },
      {
        name: 'About Me',
        columns: [
          {
            size: 'col-md',
            content: {
              type: 'text',
              body: `
              <p>
                Lorem ipsum dolor sit amet consectetur adipisicing elit. Officia alias distinctio praesentium similique sit
                quaerat neque qui placeat et, possimus ut veniam laborum earum? Laborum eaque doloribus aperiam nesciunt
                repellat!
              </p>
              <p>
                Lorem ipsum dolor sit amet consectetur adipisicing elit. Officia alias distinctio praesentium similique sit
                quaerat neque qui placeat et, possimus ut veniam laborum earum? Laborum eaque doloribus aperiam nesciunt
                repellat!
              </p>
              <p>
                Lorem ipsum dolor sit amet consectetur adipisicing elit. Officia alias distinctio praesentium similique sit
                quaerat neque qui placeat et, possimus ut veniam laborum earum? Laborum eaque doloribus aperiam nesciunt
                repellat!
              </p>
            `
            }
          },
          {
            size: 'col-md',
            content: {
              type: 'progressbar',
              progressBars: [
                {
                  name: 'Java',
                  value: 70
                },
                {
                  name: 'ReactJs',
                  value: 703
                },
                {
                  name: 'Angular',
                  value: 85
                },
                {
                  name: 'Laravel',
                  value: 80
                },
                {
                  name: 'Git',
                  value: 90
                },
                {
                  name: 'NodeJs',
                  value: 705
                },
              ]
            }
          },
          {
            size: 'col-12 pb-3 pt-3',
            content: {
              type: 'text',
              body: `
              <div class="text-center">
                <a class="btn btn-primary" href="#" role="button">Download Resume</a>
              </div>
            `
            }
          }
        ]
      },
      {
        name: 'Posts',
        posts: [...Array(6).keys()].map(value => {
          return {
            title: faker.lorem.sentence(faker.random.number({ min: 3, max: 10 })),
            created_at: faker.date.recent(),
            created_by: faker.name.findName(),
            excerpt: faker.lorem.sentences(faker.random.number({ min: 2, max: 10 })),
            coverImage: {
              src: `https://picsum.photos/500/320?random=${faker.random.number(1080)}`,
              caption: faker.lorem.sentences(faker.random.number({ min: 1, max: 5 }))
            }
          };
        }).sort((a, b) => b.created_at.getTime() - a.created_at.getTime())
      },
      {
        name: 'Projects',
        projects: [...Array(3).keys()].map(value => {
          return {
            name: faker.lorem.sentence(faker.random.number({ min: 3, max: 10 })),
            description: faker.lorem.paragraph(),
            created_at: faker.date.recent(),
            images: [
              {
                src: `https://picsum.photos/980/600?random=${faker.random.number(1080)}`,
                caption: faker.lorem.sentence
              },
              {
                src: `https://picsum.photos/980/600?random=${faker.random.number(1080)}`,
                caption: faker.lorem.sentence
              },
              {
                src: `https://picsum.photos/980/600?random=${faker.random.number(1080)}`,
                caption: faker.lorem.sentence
              }
            ]
          };
        })
      }
    ]
  }];

  constructor() { }

  findPage(path: string): Observable<any> {
    const page = this.pages.find(page => {
      return page.path === path;
    });

    return of(page);
  }

  getLinks(): Observable<ILink[]> {
    return of(this.pages.map(page => {
      return {
        url: page.path,
        label: page.name
      };
    }));
  }
}
