CREATE TABLE IF NOT EXISTS users (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  provider VARCHAR(32) NOT NULL,
  subject VARCHAR(191) NOT NULL,
  email VARCHAR(191) NOT NULL,
  name VARCHAR(191) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY uq_provider_subject (provider, subject)
);

CREATE TABLE IF NOT EXISTS stay_cards (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  user_id BIGINT NOT NULL,
  token VARCHAR(64) NOT NULL UNIQUE,
  accommodation_name VARCHAR(191) NOT NULL,
  guest_name VARCHAR(191) NULL,
  subtitle VARCHAR(191) NULL,
  check_in_at DATETIME NOT NULL,
  check_out_at DATETIME NOT NULL,
  valid_from DATETIME NOT NULL,
  valid_until DATETIME NOT NULL,
  delete_after DATETIME NOT NULL,
  address TEXT NOT NULL,
  maps_url TEXT NULL,
  entry_type VARCHAR(100) NULL,
  entry_instructions TEXT NULL,
  keybox_code VARCHAR(100) NULL,
  wifi_ssid VARCHAR(191) NULL,
  wifi_password VARCHAR(191) NULL,
  house_info TEXT NULL,
  contact_name VARCHAR(191) NULL,
  contact_phone VARCHAR(100) NULL,
  contact_whatsapp VARCHAR(100) NULL,
  is_active BOOLEAN NOT NULL DEFAULT TRUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT fk_card_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
