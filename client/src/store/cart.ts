import { create } from "zustand";
import { persist, createJSONStorage } from "zustand/middleware";

type BearState = {
  books: Record<string, Book>;
  addBook: (book: Book) => void;
  updateQuantity: (
    id: string,
    quantity: number,
    options?: { onSuccess?: () => void },
  ) => void;
  removeBook: (id: string) => void;
};

type Book = {
  id: string;
  title: string;
  quantity: number;
  price: number;
};

export const useCartStore = create<BearState>()(
  persist(
    (set, get) => ({
      books: {},
      addBook: (book: Book) => {
        if (book.id in get().books) {
          console.log("cannot add book: already in the store");
          return;
        }
        set({ books: { ...get().books, [book.id]: book } });
      },
      removeBook: (id: string) => {
        const books = { ...get().books };
        if (id in books) {
          delete books[id];
          set({ books: books });
          return;
        }
        console.log("cannot delete book: no such id in store");
      },

      updateQuantity: (id: string, quantity: number, { onSuccess } = {}) => {
        const books = get().books;
        if (!(id in books)) {
          console.log("cannot update quantity: no such id in store");
          return;
        }

        books[id] = { ...books[id]!, quantity };
        onSuccess?.();
        set({ books });
      },
    }),
    {
      name: "cart",
      storage: createJSONStorage(() => sessionStorage),
    },
  ),
);
