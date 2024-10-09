import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { StrictMode, Suspense } from "react";
import { createRoot } from "react-dom/client";
import "./styles/reset.css";
import "./styles/index.css";
import "@mantine/core/styles.css";
import "@mantine/notifications/styles.css";
import { queryClient } from "./lib/tanstackQuery";
import { QueryClientProvider } from "@tanstack/react-query";
import { createTheme, Loader, MantineProvider } from "@mantine/core";
import { Notifications } from "@mantine/notifications";
import { BooksPage } from "./pages";
import { MainLayout } from "./layouts";

export const theme = createTheme({
  /* Put your mantine theme override here */
});

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <QueryClientProvider client={queryClient}>
      <MantineProvider theme={theme}>
        <Notifications position="top-right" />
        <Suspense fallback={<Loader />}>
          <MainLayout>
            <BooksPage />
          </MainLayout>
        </Suspense>
      </MantineProvider>
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  </StrictMode>,
);
