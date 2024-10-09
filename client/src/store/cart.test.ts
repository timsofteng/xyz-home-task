import { beforeEach, describe, it, expect } from "vitest";
import { useCartStore } from "./cart";

const mockBook = { id: "1", title: "Test Book", quantity: 1, price: 44.29 };

beforeEach(() => {
  // Reset Zustand store state before each test
  useCartStore.setState({ books: {} });
});

describe("useCartStore", () => {
  it("should add a book to the cart", () => {
    const { addBook } = useCartStore.getState();
    addBook(mockBook);

    const { books } = useCartStore.getState();
    expect(books["1"]).toEqual(mockBook);
  });

  it("should not add a book if it is already in the cart", () => {
    const { addBook } = useCartStore.getState();

    addBook(mockBook);
    addBook(mockBook);

    const { books } = useCartStore.getState();

    expect(Object.keys(books)).toHaveLength(1);
  });

  it("should remove a book from the cart", () => {
    const { addBook, removeBook } = useCartStore.getState();

    addBook(mockBook);
    removeBook(mockBook.id);

    const { books } = useCartStore.getState();
    expect(books[mockBook.id]).toBeUndefined();
  });

  it("should update the quantity of a book in the cart", () => {
    const { addBook, updateQuantity } = useCartStore.getState();

    addBook(mockBook);
    updateQuantity("1", 3);

    const { books } = useCartStore.getState();

    expect(books["1"]?.quantity).toBe(3);
  });

  it("should not update quantity for a non-existent book", () => {
    const { updateQuantity } = useCartStore.getState();

    updateQuantity("2", 3);

    expect(useCartStore.getState().books["2"]).toBeUndefined();
  });
});
