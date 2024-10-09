import { useMemo } from "react";
import { useCartStore } from "@store/cart";
import { CartItem } from "./CartItem";
import styled from "styled-components";

export const CartList = () => {
  const { books } = useCartStore();

  const booksArr = Object.values(books);

  const totalPrice = useMemo(() => {
    let price = 0;
    booksArr.forEach((b) => {
      price += b.price * b.quantity;
    });

    return price.toFixed(2);
  }, [booksArr]);

  return (
    <SWrapper>
      {booksArr.length < 1
        ? "No items in cart"
        : booksArr.map((b) => (
            <div key={b.id}>
              <CartItem {...b} />
            </div>
          ))}
      {booksArr.length > 0 && <span>Total price: {totalPrice}</span>}
    </SWrapper>
  );
};

const SWrapper = styled.section`
  position: sticky;
  top: 96px;

  overflow: auto;
  display: grid;
  gap: var(--space);
  align-content: baseline;

  min-width: 380px;
  height: fit-content;
  max-height: calc(100vh - 200px);
  padding: var(--space);

  border: 1px solid;
`;
