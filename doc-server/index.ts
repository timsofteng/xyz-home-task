import { serve } from "bun";
import { join } from "path";

// Путь к статическим файлам
const staticDir = join(import.meta.dir, "static");

serve({
  port: 8080,
  fetch(req) {
    const url = new URL(req.url);

    // Проверяем, запрашивается ли HTML-документ
    if (url.pathname === "/" || url.pathname === "/redoc") {
      return new Response(Bun.file(join(staticDir, "index.html")));
    }

    // Обработка других маршрутов
    return new Response("Not Found", { status: 404 });
  },
});

console.log("Server is running on http://localhost:8080");
