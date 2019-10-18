export interface Section {
  type: string;
  deleted?: boolean;
  updated?: boolean;
  order: number;
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
