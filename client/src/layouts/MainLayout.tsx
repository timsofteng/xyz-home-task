import { ErrorBoundary } from "react-error-boundary";
import styled from "styled-components";
import { useCartStore } from "@store/cart";
import { ReactNode, Suspense, useState } from "react";
import { Button, Loader } from "@mantine/core";
import { CartList } from "@components/cart";

type Props = { children: ReactNode };

export const MainLayout = ({ children }: Props) => {
  const [cartOpen, setCartOpen] = useState(false);

  return (
    <>
      <Header onCartClick={() => setCartOpen((state) => !state)} />
      <SContent>
        <ErrorBoundary fallback={<div>Something went wrong</div>}>
          <Suspense fallback={<Loader />}>
            {children}
            {cartOpen && <CartList />}
          </Suspense>
        </ErrorBoundary>
      </SContent>
    </>
  );
};

export const Header = ({ onCartClick }: { onCartClick: () => void }) => {
  const count = useCartStore(({ books }) => {
    let res = 0;
    Object.values(books).forEach((b) => {
      res += b.quantity;
    });

    return res;
  });
  return (
    <SHeader>
      <span>XYZ Bookshop</span>
      <Button onClick={onCartClick}>Cart {count || ""}</Button>
    </SHeader>
  );
};

const SHeader = styled.header`
  position: sticky;
  z-index: 2;
  top: 0;

  display: flex;
  align-items: center;
  justify-content: space-between;

  margin-bottom: var(--space);
  padding: var(--space-s);

  font-weight: 600;
  color: var(--contrast-color);

  background: var(--contrast-bg);
`;

const SContent = styled.section`
  display: flex;
  gap: var(--space-l);
  padding: 1em;
`;
