export interface Section {
  type: string;
}

export interface Link {
  text: string;
  link: string;
}

export interface HeroSection extends Section {
  image: string;
  links: {
    text: string;
    link: string;
  }[];
  content: string;
}
