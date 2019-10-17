import { Section } from './section';
export interface Page {
  name: string;
  path: string;
  description: string;
  sections: Section[];
}
