export interface Post {
  id?: string;
  title: string;
  created_at: Date;
  created_by: string;
  excerpt?: string;
  coverImage?: {
    src: string,
    caption?: string
  };
  body?: string;
}
