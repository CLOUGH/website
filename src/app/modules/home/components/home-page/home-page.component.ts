import { Component, OnInit } from '@angular/core';
import * as faker from 'faker';
import { Post } from '../../models/post';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.scss']
})
export class HomePageComponent implements OnInit {

  pageContent = {
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
  };

  posts: Post[] = [];

  constructor() { }

  ngOnInit() {
    const noPost = 6;
    for (let i = 0; i < noPost; i++) {
      this.posts.push(
        {
          title: faker.lorem.sentence(faker.random.number({ min: 3, max: 10 })),
          created_at: faker.date.recent(),
          created_by: faker.name.findName(),
          excerpt: faker.lorem.sentences(faker.random.number({ min: 2, max: 10 })),
          coverImage: {
            src: `https://picsum.photos/500/320?random=${faker.random.number(1080)}`,
            caption: faker.lorem.sentences(faker.random.number({ min: 1, max: 5 }))
          }
        }
      );
    }

    this.posts.sort((a, b) => b.created_at.getTime() - a.created_at.getTime());
  }
}
