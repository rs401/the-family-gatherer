
export interface User {
  ID: number;
  Name: string
}
export interface Forum {
    ID: number;
    Name: string;
    UserId: number;
    User: User;
}

export interface Thread {
  ID: number;
  Title: string;
  Body: string;
  UserId: number;
  User: User;
  ForumId: number;
  Forum: Forum;
  Posts: Post[];
}

export interface Post {
  ID: number;
  Body: string;
  UserId: number;
  User: User;
  ThreadId: number;
  Thread: Thread;
}