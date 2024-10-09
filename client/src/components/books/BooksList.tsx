import { ReactNode } from "react";
import styled from "styled-components";

type Props = {
  children: ReactNode;
};

export const BooksList = ({ children }: Props) => {
  return <SList>{children}</SList>;
};

const SList = styled.section`
  display: grid;
  gap: var(--space-l);
`;
