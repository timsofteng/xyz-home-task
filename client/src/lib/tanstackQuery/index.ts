import { QueryClient } from "@tanstack/react-query";

export const queryClient = new QueryClient({
  defaultOptions: { queries: { gcTime: Infinity, staleTime: Infinity } },
});
