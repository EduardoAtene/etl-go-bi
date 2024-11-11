-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: emcash_db
-- Tempo de geração: 11/11/2024 às 04:58
-- Versão do servidor: 10.4.34-MariaDB-1:10.4.34+maria~ubu2004
-- Versão do PHP: 8.2.8

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Banco de dados: `DW_Acidentes`
--

-- --------------------------------------------------------

--
-- Estrutura para tabela `Dim_Condicoes`
--

CREATE TABLE `Dim_Condicoes` (
  `id_condicoes` int(11) NOT NULL,
  `condicao_meteorologica` varchar(50) DEFAULT NULL,
  `tipo_pista` varchar(50) DEFAULT NULL,
  `tracado_via` varchar(100) DEFAULT NULL,
  `uso_solo` varchar(256) DEFAULT NULL,
  `sentido_via` varchar(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

-- --------------------------------------------------------

--
-- Estrutura para tabela `Dim_Localizacao`
--

CREATE TABLE `Dim_Localizacao` (
  `id_localizacao` int(11) NOT NULL,
  `municipio` varchar(100) DEFAULT NULL,
  `br` varchar(256) DEFAULT NULL,
  `km` int(11) DEFAULT NULL,
  `latitude` decimal(10,8) DEFAULT NULL,
  `longitude` decimal(11,8) DEFAULT NULL,
  `regional` varchar(50) DEFAULT NULL,
  `delegacia` varchar(50) DEFAULT NULL,
  `uop` varchar(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

-- --------------------------------------------------------

--
-- Estrutura para tabela `Dim_Pessoa`
--

CREATE TABLE `Dim_Pessoa` (
  `id_pessoa` int(11) NOT NULL,
  `tipo_envolvido` varchar(50) DEFAULT NULL,
  `idade` int(11) DEFAULT NULL,
  `sexo` varchar(20) DEFAULT NULL,
  `raca_cor` varchar(50) DEFAULT NULL,
  `estado_fisico` varchar(50) DEFAULT NULL,
  `municipio_residencia` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

-- --------------------------------------------------------

--
-- Estrutura para tabela `Dim_Tempo`
--

CREATE TABLE `Dim_Tempo` (
  `id_tempo` int(11) NOT NULL,
  `data_completa` date DEFAULT NULL,
  `ano` int(11) DEFAULT NULL,
  `mes` int(11) DEFAULT NULL,
  `dia` int(11) DEFAULT NULL,
  `dia_semana` varchar(20) DEFAULT NULL,
  `hora` time DEFAULT NULL,
  `periodo_dia` varchar(256) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

-- --------------------------------------------------------

--
-- Estrutura para tabela `Dim_Veiculo`
--

CREATE TABLE `Dim_Veiculo` (
  `id_veiculo` int(11) NOT NULL,
  `tipo_veiculo` varchar(50) DEFAULT NULL,
  `marca` varchar(100) DEFAULT NULL,
  `ano_fabricacao` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

-- --------------------------------------------------------

--
-- Estrutura para tabela `Fato_Acidentes`
--

CREATE TABLE `Fato_Acidentes` (
  `id_acidente` int(11) NOT NULL,
  `id_tempo` int(11) DEFAULT NULL,
  `id_localizacao` int(11) DEFAULT NULL,
  `id_veiculo` int(11) DEFAULT NULL,
  `id_pessoa` int(11) DEFAULT NULL,
  `id_condicoes` int(11) DEFAULT NULL,
  `fonte_dados` varchar(10) DEFAULT NULL,
  `causa_acidente` varchar(100) DEFAULT NULL,
  `tipo_acidente` varchar(100) DEFAULT NULL,
  `classificacao_acidente` varchar(50) DEFAULT NULL,
  `cid_causa_morte` varchar(10) DEFAULT NULL,
  `desc_causa_morte` varchar(200) DEFAULT NULL,
  `qtd_ilesos` int(11) DEFAULT NULL,
  `qtd_feridos_leves` int(11) DEFAULT NULL,
  `qtd_feridos_graves` int(11) DEFAULT NULL,
  `qtd_mortos` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Índices para tabelas despejadas
--

--
-- Índices de tabela `Dim_Condicoes`
--
ALTER TABLE `Dim_Condicoes`
  ADD PRIMARY KEY (`id_condicoes`);

--
-- Índices de tabela `Dim_Localizacao`
--
ALTER TABLE `Dim_Localizacao`
  ADD PRIMARY KEY (`id_localizacao`);

--
-- Índices de tabela `Dim_Pessoa`
--
ALTER TABLE `Dim_Pessoa`
  ADD PRIMARY KEY (`id_pessoa`);

--
-- Índices de tabela `Dim_Tempo`
--
ALTER TABLE `Dim_Tempo`
  ADD PRIMARY KEY (`id_tempo`);

--
-- Índices de tabela `Dim_Veiculo`
--
ALTER TABLE `Dim_Veiculo`
  ADD PRIMARY KEY (`id_veiculo`);

--
-- Índices de tabela `Fato_Acidentes`
--
ALTER TABLE `Fato_Acidentes`
  ADD PRIMARY KEY (`id_acidente`),
  ADD KEY `idx_fato_tempo` (`id_tempo`),
  ADD KEY `idx_fato_localizacao` (`id_localizacao`),
  ADD KEY `idx_fato_veiculo` (`id_veiculo`),
  ADD KEY `idx_fato_pessoa` (`id_pessoa`),
  ADD KEY `idx_fato_condicoes` (`id_condicoes`);

--
-- AUTO_INCREMENT para tabelas despejadas
--

--
-- AUTO_INCREMENT de tabela `Dim_Condicoes`
--
ALTER TABLE `Dim_Condicoes`
  MODIFY `id_condicoes` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `Dim_Localizacao`
--
ALTER TABLE `Dim_Localizacao`
  MODIFY `id_localizacao` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `Dim_Pessoa`
--
ALTER TABLE `Dim_Pessoa`
  MODIFY `id_pessoa` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `Dim_Tempo`
--
ALTER TABLE `Dim_Tempo`
  MODIFY `id_tempo` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `Dim_Veiculo`
--
ALTER TABLE `Dim_Veiculo`
  MODIFY `id_veiculo` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de tabela `Fato_Acidentes`
--
ALTER TABLE `Fato_Acidentes`
  MODIFY `id_acidente` int(11) NOT NULL AUTO_INCREMENT;

--
-- Restrições para tabelas despejadas
--

--
-- Restrições para tabelas `Fato_Acidentes`
--
ALTER TABLE `Fato_Acidentes`
  ADD CONSTRAINT `Fato_Acidentes_ibfk_1` FOREIGN KEY (`id_tempo`) REFERENCES `Dim_Tempo` (`id_tempo`),
  ADD CONSTRAINT `Fato_Acidentes_ibfk_2` FOREIGN KEY (`id_localizacao`) REFERENCES `Dim_Localizacao` (`id_localizacao`),
  ADD CONSTRAINT `Fato_Acidentes_ibfk_3` FOREIGN KEY (`id_veiculo`) REFERENCES `Dim_Veiculo` (`id_veiculo`),
  ADD CONSTRAINT `Fato_Acidentes_ibfk_4` FOREIGN KEY (`id_pessoa`) REFERENCES `Dim_Pessoa` (`id_pessoa`),
  ADD CONSTRAINT `Fato_Acidentes_ibfk_5` FOREIGN KEY (`id_condicoes`) REFERENCES `Dim_Condicoes` (`id_condicoes`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
