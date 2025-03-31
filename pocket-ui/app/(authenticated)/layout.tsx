import type { Metadata } from "next";
import { redirect } from 'next/navigation';

import { auth } from "@/app/lib/auth";
import ClientSideLayout from './ClientSideLayout';

export const metadata: Metadata = {
  title: "e-pocket",
  description: "streamlined e-wallet experience",
};

export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const session = await auth();
  if (session == null) {
    return redirect("/");
  }

  return (
    <ClientSideLayout>
      {children}
    </ClientSideLayout>
  );
}
