-- Inserts

-- Pilotos
INSERT INTO pilotos (
    id, nome, email, iracing_id, irating, imagem, created_by, created_at, updated_at
) VALUES
(1, 'João Silva', 'joao.silva@email.com', 1234567, 2500, 'joao.png', 'admin', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 'Maria Souza', 'maria.souza@email.com', 2345678, 3200, 'maria.png', 'admin', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 'Max Verstappen', 'max.v@email.com', 3456789, 8000, 'max.png', 'admin', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(4, 'Tony Kanaan', 'tony.kanaan@email.com', 4567890, 6500, 'tony.png', 'admin', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(5, 'Rubens Barrichello', 'rubens.b@email.com', 5678901, 6000, 'rubens.png', 'admin', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(6, 'Felipe Massa', 'felipe.m@email.com', 6789012, 5800, 'massa.png', 'admin', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(7, 'Bruno Senna', 'bruno.senna@email.com', 7890123, 5700, 'bruno.png', 'admin', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(8, 'Enzo Fittipaldi', 'enzo.f@email.com', 8901234, 5400, 'enzo.png', 'admin', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(9, 'Gabriel Casagrande', 'gabriel.c@email.com', 9012345, 5100, 'gabriel.png', 'admin', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(10, 'Augusto Farfus', 'augusto.f@email.com', 1123456, 6900, 'farfus.png', 'admin', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Carros
INSERT INTO carros (id, nome, classe, imagem, created_at, updated_at) VALUES
(1, 'Porsche 911 GT3', 'GT3', 'porsche.png', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 'BMW M4 GT3', 'GT3', 'bmw.png', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 'Audi R8 LMS', 'GT3', 'audi.png', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(4, 'Mercedes AMG GT3', 'GT3', 'mercedes.png', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(5, 'Ferrari 488 GTE', 'GTE', 'ferrari.png', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(6, 'Corvette C8.R', 'GTE', 'corvette.png', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(7, 'Dallara P217', 'LMP2', 'dallara.png', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(8, 'Mazda RT24-P', 'DPi', 'mazda.png', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(9, 'Lamborghini Huracan GT3', 'GT3', 'lamborghini.png', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(10, 'Ford GT', 'GTE', 'fordgt.png', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- PilotoCarrosDisponiveis
INSERT INTO piloto_carros_disponiveis (piloto_id, carro_id) VALUES
(1, 1), (1, 2),
(2, 2), (2, 3),
(3, 4), (3, 7),
(4, 5), (4, 6),
(5, 7), (5, 8),
(6, 3), (6, 9),
(7, 1), (7, 10),
(8, 5), (8, 2),
(9, 6), (9, 4),
(10, 8), (10, 9);

-- Pista (singular)
INSERT INTO pista (id, nome, localizacao, imagem, created_at, updated_at) VALUES
(1, 'Spa-Francorchamps', 'Bélgica', 'spa.png', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 'Monza', 'Itália', 'monza.png', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 'Le Mans', 'França', 'lemans.png', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(4, 'Sebring', 'EUA', 'sebring.png', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(5, 'Daytona', 'EUA', 'daytona.png', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Eventos
INSERT INTO eventos (
    id, nome, plataforma, data_evento, pista_id, duracao, classes, min_pilotos, max_pilotos, imagem, created_by, created_at, updated_at
) VALUES
(1, 'Endurance 6h Spa', 'iRacing', '2025-12-10', 1, 360, ARRAY['GT3','GTE','LMP2'], 2, 5, 'spa_event.png', 'admin', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 'Sprint 1h Monza', 'iRacing', '2025-12-17', 2, 60, ARRAY['GT3'], 1, 3, 'monza_event.png', 'admin', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 'Le Mans 24h', 'iRacing', '2026-01-15', 3, 1440, ARRAY['GT3','GTE','LMP2','DPi'], 3, 6, 'lemans_event.png', 'admin', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(4, 'Sebring 12h', 'iRacing', '2026-02-20', 4, 720, ARRAY['GTE','LMP2'], 2, 5, 'sebring_event.png', 'admin', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(5, 'Daytona 2.4h', 'iRacing', '2026-03-05', 5, 144, ARRAY['GT3','DPi'], 2, 4, 'daytona_event.png', 'admin', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);