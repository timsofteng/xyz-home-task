import { BookItem, BooksList } from "@components/books";
import { TextInput } from "@mantine/core";
import { useState } from "react";
import { $api } from "@lib/api";

export const BooksPage = () => {
  const [query, setQuery] = useState("go");
  const {
    data: { items },
  } = $api.useSuspenseQuery("get", "/books", {
    params: {
      query: { q: query || "go" },
    },
  });

  return (
    <BooksList>
      <TextInput
        placeholder="Filter books..."
        onKeyUpCapture={({ key, currentTarget: { value } }) => {
          if (key === "Enter") {
            setQuery(value);
          }
        }}
      />
      <h2>Book results</h2>
      {items.length > 0
        ? items.map((i) => <BookItem key={i.id} {...i} />)
        : "No books"}
    </BooksList>
  );
};
