export const metadata = {
  title: '{{PROJECT_NAME}}',
  description: 'Generated with LutzKit',
}

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  )
}
