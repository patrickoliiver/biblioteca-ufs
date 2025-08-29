// Script para testar conectividade com PostgreSQL da Amazon
const { Client } = require('pg');

const connectionString = 'postgresql://principal:senha-principa@database-postgre-bd1-2025-1.cdwk8iisko7f.us-east-1.rds.amazonaws.com:5432/postgres?sslmode=require';

async function testConnection() {
  const client = new Client({
    connectionString: connectionString,
    ssl: {
      rejectUnauthorized: false
    }
  });

  try {
    console.log('🔌 Testando conexão com PostgreSQL da Amazon...');
    await client.connect();
    console.log('✅ Conectado com sucesso!');
    
    // Teste de query
    const result = await client.query('SELECT NOW() as current_time');
    console.log('⏰ Hora atual do banco:', result.rows[0].current_time);
    
    // Verificar tabelas
    const tables = await client.query(`
      SELECT table_name 
      FROM information_schema.tables 
      WHERE table_schema = 'public'
    `);
    console.log('📋 Tabelas encontradas:', tables.rows.map(r => r.table_name));
    
  } catch (error) {
    console.error('❌ Erro na conexão:', error.message);
  } finally {
    await client.end();
  }
}

testConnection();
