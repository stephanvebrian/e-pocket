import type { Metadata } from "next";
import { cookies } from 'next/headers';

import "./globals.css";

export const metadata: Metadata = {
  title: "e-pocket",
  description: "revolution of digital wallet",
};

export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const cookieStore = await cookies();
  const theme = cookieStore.get('theme')?.value || "light";

  return (
    <html lang="en" data-theme={theme}>
      <body>
        {children}
      </body>
    </html >
  );
}
