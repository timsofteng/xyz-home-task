import { Button } from "@mantine/core";
import { useCartStore } from "@store/cart";
import styled from "styled-components";
import { notifications } from "@mantine/notifications";
import { Form, useForm } from "react-hook-form";
import { z } from "zod";
import { NumberInput } from "react-hook-form-mantine";
import { zodResolver } from "@hookform/resolvers/zod";

type Props = {
  id: string;
  title: string;
  quantity: number;
};

const CUSTOM_ERROR_MSG = "Sorry! Number invalid";

const schema = z.object({
  quantity: z
    .number()
    .min(1, { message: CUSTOM_ERROR_MSG })
    .max(100, { message: CUSTOM_ERROR_MSG }),
});

type FormSchemaType = z.infer<typeof schema>;

export const CartItem = ({ id, quantity, title }: Props) => {
  const { removeBook, updateQuantity } = useCartStore();
  const { control } = useForm<FormSchemaType>({
    mode: "onChange",
    defaultValues: { quantity },
    resolver: zodResolver(schema),
  });

  return (
    <Form
      control={control}
      onError={(e) => console.log(e)}
      onSubmit={({ data }) => {
        updateQuantity(id, data.quantity, {
          onSuccess: () => {
            notifications.show({ message: "Quantity has been updated" });
          },
        });
      }}
    >
      <SWrapper>
        <SHeader>
          <h5>{title}</h5>
          <Button
            size="xs"
            style={{ background: "red" }}
            onClick={() => removeBook(id)}
          >
            x
          </Button>
        </SHeader>
        <NumberInput min={1} control={control} name="quantity" />
        <Button type="submit">Update quantity</Button>
      </SWrapper>
    </Form>
  );
};

const SWrapper = styled.div`
  position: relative;

  display: grid;
  gap: var(--space);

  padding: var(--space-s);

  border: 1px solid;
`;

const SHeader = styled.div`
  display: flex;
  gap: var(--space);
  align-items: center;
  justify-content: space-between;
`;
