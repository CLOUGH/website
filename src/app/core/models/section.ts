export interface Section {
  type: string;
}

export interface Link {
  text: string;
  link: string;
}

export interface HeroSection extends Section {
  image: string;
  content: string;
}

export interface TextSection extends Section {
  content: string;
}
