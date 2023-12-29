interface User {
  name: string;
  checkoutLimit: number;
  checkedOut: BookCopy[];
}

interface BookCopy {
  holder: User;
  checkoutDate: Date;
  dueDate: Date;
  book: Book;
  slot: BookSlot;
  userPosition: number;
  library: Library;

  // reads from where last left off
  read(): string;
}

interface Book {
  title: string;
  author: string;
  pageCount: number;

  // reads from page 1 or page specified (userPosition)
  read(page?: number): string;
}

interface BookSlot {
  copies: number;
  book: Book;

  // returns a book copy if available and user is not over checkout limit
  checkout(user: User): void;
  // increments number of copies and removes book copy from user
  returnBook(): void;
}

interface Library {
  books: BookSlot[];
  users: User[];

  // creates a new book slot if book does not exist
  addBook(book: Book, copies?: number): void;
  // removes book slot if no copies left
  removeBook(book: Book, copies?: number): void;
  addUser(user: User): void;
  removeUser(user: User): void;
  // finds bookslot and checks out book copy to user
  checkoutBook(book: Book, user: User): void;
  // finds bookslot and returns book copy from user
  returnBook(bookCopy: BookCopy): void;
}
