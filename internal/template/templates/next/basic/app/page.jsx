export default function Home() {
  return (
    <main style={{
      minHeight: '100vh',
      display: 'flex',
      flexDirection: 'column',
      alignItems: 'center',
      justifyContent: 'center',
      fontFamily: 'system-ui',
      textAlign: 'center',
    }}>
      <h1>Hello from {{PROJECT_NAME}}!</h1>
      <p>Your Next.js project is ready.</p>
    </main>
  )
}
