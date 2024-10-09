import { Button, Divider, Text } from "@mantine/core";
import { styled } from "styled-components";
import { useCartStore } from "@store/cart";

type Props = {
  id: string;
  thumbnail: string;
  title: string;
  description: string;
  price: number;
  currency: string;
  pages: number;
};

export const BookItem = ({
  id,
  thumbnail,
  description,
  currency,
  price,
  title,
  pages,
}: Props) => {
  const { addBook } = useCartStore();
  return (
    <SWrapper>
      <SThumbnail>
        {thumbnail.length ? (
          <img height={218} width={150} src={thumbnail} />
        ) : (
          "No image"
        )}
      </SThumbnail>
      <SContent>
        <b>{title}</b>
        {description.length ? (
          <div>{description}</div>
        ) : (
          <Text c="dimmed">No description</Text>
        )}
        <div>Pages: {pages || "No info"}</div>
        <div>Price: {price === 0 ? "Free" : `${price} ${currency}`}</div>
        <Button onClick={() => addBook({ id, quantity: 1, title, price })}>
          Add to cart
        </Button>
      </SContent>
      <Divider my="md" />
    </SWrapper>
  );
};

const SWrapper = styled.article`
  display: flex;
  gap: var(--space);
`;

const SContent = styled.div`
  display: grid;
  flex-grow: 1;
  gap: var(--space-s);
`;

const SThumbnail = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;

  min-width: 150px;
  height: 220px;

  border: 1px solid;
`;
